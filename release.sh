#!/bin/bash

# Variables
# ---------
CURRENT_DIR="$(dirname "$0")"
APP_NAME="goRestBoilerplate"

# Arguments
# ---------
if [ $# -ne 1 ]; then
  echo "Usage: ./release.sh [version (x.y.z)]"
  exit 1
fi
VERSION=$1
echo "Version: $VERSION"
echo

# 1. Création du répertoire
# -------------------------
echo "1. Création du répertoire $VERSION..."
mkdir -p releases/$VERSION
echo "   => Répertoire releases/$VERSION créé avec succès"
RELEASE_DIR=releases/$VERSION

# 2. Build du binaire
# -------------------
RELEASE_NAME=$APP_NAME"_"$VERSION
echo "2. Build..."
go build -o $RELEASE_NAME
if [ $? -ne 0 ]; then
  echo "[Erreur] Echec du build"

  rm -rf $RELEASE_DIR
  exit 2
fi
echo "   => Build créée avec succès"

# 3. Déplacement de la release et di fichier de configuration
# -----------------------------------------------------------
echo "3. Déplacement des fichiers..."
mv $RELEASE_NAME $RELEASE_DIR
cp config.toml $RELEASE_DIR
echo "   => Déplacement des fichiers réalisé avec succès"

# 4. Renommage du numéro de version dans le fichier de configuration
# ------------------------------------------------------------------
echo "4. Renommage du numéro de version dans le fichier de configuration..."
# sed -i 's//g' RELEASE_DIR/config.toml
sed -i -e 's/^\(version\s*=\s*\).*$/\1$VERSION/' $RELEASE_DIR/config.toml
