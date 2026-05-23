#!/bin/bash

# Install lsgo as an alias into the bash profile

# Create ~/.local/bin if it does not exists
mkdir -p ~/.local/bin

# In ~/.profile or .bashrc or whichever one is the most appropriate one
# Add ~/.local/bin to PATH
PROFILE_FILE="$HOME/.bashrc"
if [ -f "$HOME/.bash_profile" ]; then
  PROFILE_FILE="$HOME/.bash_profile"
elif [ -f "$HOME/.profile" ]; then
  PROFILE_FILE="$HOME/.profile"
fi

if ! grep -q '$HOME/.local/bin' "$PROFILE_FILE"; then
  echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$PROFILE_FILE"
fi

# And then finally add the alias: `alias ll=lsgo`
if ! grep -q 'alias ll=lsgo' "$PROFILE_FILE"; then
  echo 'alias ll=lsgo' >> "$PROFILE_FILE"
fi

echo "Alias ll=lsgo installed. Please run 'source $PROFILE_FILE' or restart your terminal."
