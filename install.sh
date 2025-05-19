#!/usr/bin/env bash
set -euo pipefail

# ─── Config ────────────────────────────────────────────────────────────────
VERSION="v1.0.1"
BASE_URL="https://github.com/HISP-Uganda/dhis2cli/releases/download/${VERSION}"
LINUX_ASSET="dhis2_${VERSION#v}_Linux_x86_64.tar.gz"
DARWIN_ASSET="dhis2_${VERSION#v}_Darwin_x86_64.tar.gz"

INSTALL_DIR="${1:-/usr/local/bin}"

# ─── Helpers ───────────────────────────────────────────────────────────────
info()    { printf "ℹ️  %s\n" "$*" >&2; }
success() { printf "✅ %s\n" "$*" >&2; }
warn()    { printf "⚠️  %s\n" "$*" >&2; }

# ─── 1. Detect OS and pick asset ────────────────────────────────────────────
OS="$(uname -s)"
case "$OS" in
  Darwin) ASSET="$DARWIN_ASSET" ;;
  Linux)  ASSET="$LINUX_ASSET"  ;;
  *)
    warn "Unrecognized OS '$OS', defaulting to Linux build."
    ASSET="$LINUX_ASSET"
    ;;
esac
DOWNLOAD_URL="${BASE_URL}/${ASSET}"
info "OS detected: $OS → fetching $ASSET"

# ─── 2. Detect shell and rc file ────────────────────────────────────────────
if [[ "${SHELL##*/}" == "zsh" ]] || [[ -n "${ZSH_VERSION-}" ]]; then
  SHELL_NAME="zsh"
  SHELLRC_FILE="${ZDOTDIR:-$HOME}/.zshrc"
  COMP_CMD="completion zsh"
else
  SHELL_NAME="bash"
  SHELLRC_FILE="$HOME/.bashrc"
  COMP_CMD="completion bash"
fi
SRC_LINE="source <($INSTALL_DIR/dhis2 ${COMP_CMD})"
info "Shell detected: $SHELL_NAME → will update $SHELLRC_FILE"

# ─── 3. Download & extract ──────────────────────────────────────────────────
TMPDIR="$(mktemp -d)"
trap 'rm -rf "$TMPDIR"' EXIT

info "Downloading dhis2cli..."
curl -fsSL "$DOWNLOAD_URL" -o "$TMPDIR/dhis2.tar.gz"
success "Downloaded to $TMPDIR/dhis2.tar.gz"

info "Extracting archive..."
tar -xzf "$TMPDIR/dhis2.tar.gz" -C "$TMPDIR"
success "Extracted into $TMPDIR"

# ─── 4. Install the binary ─────────────────────────────────────────────────
BINARY_NAME="dhis2"
info "Installing '$BINARY_NAME' to '$INSTALL_DIR'..."
mkdir -p "$INSTALL_DIR"
cp "$TMPDIR/$BINARY_NAME" "$INSTALL_DIR/"
chmod +x "$INSTALL_DIR/$BINARY_NAME"
success "Installed to $INSTALL_DIR/$BINARY_NAME"

# ─── 5. Create default config ───────────────────────────────────────────────
CONFIG_FILE="$HOME/.dhis2cli.yaml"
if [[ -f "$CONFIG_FILE" ]]; then
  warn "Config file $CONFIG_FILE already exists; skipping creation."
else
  info "Creating default config at $CONFIG_FILE..."
  cat > "$CONFIG_FILE" <<EOF
server:
  database_uri: ""
  base_url: "https://play.im.dhis2.org/dev/api/"
  username: "admin"
  password: "district"
  auth_method: "Basic"
  auth_token: "123456"
EOF
  success "Config written to $CONFIG_FILE"
fi

# ─── 6. Register completion ─────────────────────────────────────────────────
info "Registering $SHELL_NAME completion in '$SHELLRC_FILE'..."
if grep -Fxq "$SRC_LINE" "$SHELLRC_FILE"; then
  success "Completion already registered."
else
  {
    printf "\n# dhis2cli %s completion\n" "$SHELL_NAME"
    printf "%s\n" "$SRC_LINE"
  } >> "$SHELLRC_FILE"
  success "Appended completion loader to $SHELLRC_FILE"
fi

# ─── 7. Done ────────────────────────────────────────────────────────────────
echo
success "Installation complete! Restart your shell or run:"
printf "  source %s\n" "$SHELLRC_FILE"
