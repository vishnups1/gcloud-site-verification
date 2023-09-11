# gcloud-site-verification

[![Build](https://github.com/vishnugth/gcloud-site-verification/actions/workflows/release.yml/badge.svg)](https://github.com/vishnugth/gcloud-site-verification/actions/workflows/release.yml)

A command-line utility for verifying and managing domain ownership using the Site Verification API.

## Pre-requisites

### Authentication

This tool uses the [Application Default Credentials](https://cloud.google.com/docs/authentication/production) strategy to authenticate with Google Cloud APIs. This means that you must have the `GOOGLE_APPLICATION_CREDENTIALS` environment variable set to the path of a valid credentials file.

**Example:**

```bash
gcloud auth application-default login --scopes "openid,https://www.googleapis.com/auth/userinfo.email,https://www.googleapis.com/auth/cloud-platform,https://www.googleapis.com/auth/sqlservice.login,https://www.googleapis.com/auth/siteverification,https://www.googleapis.com/auth/accounts.reauth"

export GOOGLE_APPLICATION_CREDENTIALS="/Users/${HOME}/.config/gcloud/application_default_credentials.json"
```

### Enable the Site Verification API

```
gcloud services enable siteverification.googleapis.com
```

## Installation

Pick the binary for your platform from the [releases page](https://github.com/vishnugth/gcloud-site-verification/releases/latest) and install it in your `$PATH`.

**Example:** `MacOS`

```bash
TEMP_DIR=$(mktemp -d)
cd $TEMP_DIR
URL="https://github.com/vishnugth/gcloud-site-verification/releases/download/v0.3.0/gcloud-site-verification_0.3.0_darwin_amd64.tar.gz"
curl -LO $URL
tar -xzvf gcloud-site-verification_0.3.0_darwin_amd64.tar.gz
chmod +x gcloud-site-verify
sudo mv gcloud-site-verify /usr/local/bin/
```

**Example:** `Linux`

```bash
TEMP_DIR=$(mktemp -d)
cd $TEMP_DIR
URL="https://github.com/vishnugth/gcloud-site-verification/releases/download/v0.3.0/gcloud-site-verification_0.3.0_linux_386.tar.gz"
curl -LO $URL
tar -xzvf gcloud-site-verification_0.3.0_linux_386.tar.gz
chmod +x gcloud-site-verify
sudo mv gcloud-site-verify /usr/local/bin/
```

## Usage

```console
gcloud-site-verify -h
A command-line utility for verifying and managing domain ownership using the Site Verification API.

Usage:
  gcloud-site-verify [command]

Available Commands:
  addowners    Adds owner(s) to a site or domain.
  getresource  Retrieves the most current data for a website or domain.
  gettoken     Gets a verification token for the authenticated user to place on a website or domain.
  help         Help about any command
  removeowners Removes owner(s) to a site or domain.
  update       Updates the list of owners for a website or domain.

Flags:
  -h, --help                help for gcloud-site-verify
  -i, --identifier string   The identifier of the site to be verified (e.g. http://www.example.com/, https://www.example.com/, dns://example.com).
  -m, --method string       The method to use for verifying a site. Check here for a list of valid methods: https://developers.google.com/site-verification/v1/getting_started#tokens (default "DNS_TXT")
  -o, --owners string       The list of owners to be added to the site. Separate multiple owners with a comma (e.g. foo@example.com,bar@example.com)
  -t, --type string         The type of site to be verified. Valid values are ANDROID_APP, INET_DOMAIN or SITE. (default "INET_DOMAIN")
  -v, --version             version for gcloud-site-verify

Use "gcloud-site-verify [command] --help" for more information about a command.
```

## Steps to verify a site

- Get a verification token.

```console
gcloud-site-verify gettoken --identifier dns://example.com --type INET_DOMAIN --method DNS_TXT
```

- Add the token as a `TXT` record in your DNS.

- Add owners.

```console
gcloud-site-verify addowners --identifier dns://example.com --owners "foo@example.com,bar@example.com"
```

- Remove owners.

```console
gcloud-site-verify removeowners --identifier dns://example.com --owners "foo@example.com"
```

- Get the list of owners for a site.

```console
gcloud-site-verify getresource -i dns://example.com
```