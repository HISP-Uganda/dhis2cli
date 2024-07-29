package users

import (
	"dhis2cli/client"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"net/http"
	"os"
)

var passwordsFile string
var password string

func init() {
	SetPasswordCmd.Flags().StringVarP(&userID, "uid", "", "", "UID for user whose password should be changed.\nNote: This is required if 'passwords-file' not provided")
	SetPasswordCmd.Flags().StringVarP(&userID, "password", "", "", "new password for user.")
	SetPasswordCmd.Flags().StringVarP(&passwordsFile, "passwords-file", "", "", `File with UIDs (one per line) for users to delete.
Note: required if 'id' flag not provided.`)
	SetPasswordCmd.MarkFlagsOneRequired("uid", "passwords-file")
	SetPasswordCmd.MarkFlagsMutuallyExclusive("uid", "passwords-file")
	SetPasswordCmd.MarkFlagsMutuallyExclusive("passwords-file", "password")
}

var SetPasswordCmd = &cobra.Command{
	Use:   "setPassword",
	Short: "Set a user's password",
	Example: `
# Set password for a single user
dhis2 users setPassword --uid <UID>

- Set password for multiple users from a CSV file of the form 
User UID,Password
UID1,password1
UID2,password2
	
- Note that the header row is ignored

# Set password for multiple users from a CSV file with a header row
dhis2 users setPassword --passwords-file users_passwords.csv
    `,
	Run: func(cmd *cobra.Command, args []string) {
		if userID != "" {
			fmt.Printf(
				`Password should be at least 8 characters long, 
with at least one lowercase character, one uppercase character, one digit and one special character.`)
			fmt.Println()
			if len(password) == 0 {
				fmt.Print("Enter Password: ")
				// password1, err := utils.ReadPasswordWithMask()
				password1 := readPassword("New Password: ")
				password2 := readPassword("Confirm New Password: ")

				if password1 != password2 {
					fmt.Println("\nPasswords do not match. Please try again.")
					os.Exit(1)
				}
				password = password1
			}

			if !utils.ValidatePassword(password) {
				fmt.Println("\nPassword does not meet the criteria. Please try again.")
				os.Exit(1)
			}
			passwordPatch := []map[string]any{
				{
					"op":    "replace",
					"path":  "/password",
					"value": password,
				},
			}
			resp, err := client.Dhis2Client.PatchResource(fmt.Sprintf("/users/%s", userID), passwordPatch)
			if err != nil {
				fmt.Printf("\nError changing user's password: %v\n", err)
				return
			}
			if resp.StatusCode() != http.StatusOK {
				fmt.Printf("\nError changing user's password (status code: %d)\n%v\n", resp.StatusCode, string(resp.Body()))
				return
			}
			fmt.Printf("\n User's password has been changed successfully.\n")
		} else {
			// using password-file now
			lines, err := utils.ReadCSV(passwordsFile)
			if err != nil {
				fmt.Printf("\nError reading passwords file: %v\n", err)
				return
			}
			for _, line := range lines {
				fmt.Println(line)
				if len(line) < 2 {
					fmt.Printf("\nInvalid line in passwords file: %v\n", line)
					continue
				}
				userID := line[0]
				password := line[1]
				if !utils.ValidatePassword(password) {
					fmt.Printf("\nPassword for user %s does not meet the criteria. Skipping...\n", userID)
					continue
				}
				passwordPatch := []map[string]any{
					{
						"op":    "replace",
						"path":  "/password",
						"value": password,
					},
				}
				resp, err := client.Dhis2Client.PatchResource(fmt.Sprintf("/users/%s", userID), passwordPatch)
				if err != nil {
					fmt.Printf("\nError changing password for user %s: %v\n", userID, err)
					continue
				}
				if resp.StatusCode() != http.StatusOK {
					fmt.Printf("\nError changing password for user %s (status code: %d)\n%v\n", userID, resp.StatusCode(), string(resp.Body()))
					continue
				}
				fmt.Printf("\nPassword for user %s has been changed successfully.\n", userID)
			}
		}
	},
}

func readPassword(prompt string) string {
	fmt.Print(prompt)
	pass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	println()
	return string(pass)
}
