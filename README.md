# HBS Algolia

A CLI Application for HBS theme.

## Installation

### Downloads

Download from [Releases](https://github.com/razonyang/hugo-theme-bootstrap-algolia/releases).

## Environments

| Name | Required | Default | Description |
|---|:-:|:-:|---|
| `ALGOLIA_APP_ID` | YES | - | The Algolia App ID.
| `ALGOLIA_API_KEY` | YES | - | The Algolia **Admin** API KEY.
| `ALGOLIA_INDEX_NAME` | YES | - | The Algolia App Index Name.
| `ALGOLIA_INDEX_FILE` | NO | `public/algolia/index.json`, relative to your working directory.

This tool requires several env vars to be set, there are multiple ways to do that.

**via command line**

```
ALGOLIA_APP_ID=x \
ALGOLIA_API_KEY=y \
ALGOLIA_INDEX_NAME=z \
hugo-theme-bootstrap-algolia 
```

**via .env file**

Create the `.env` file under your site root with following form.

```
ALGOLIA_APP_ID=x
ALGOLIA_API_KEY=y
ALGOLIA_INDEX_NAME=z
#ALGOLIA_INDEX_FILE=public/algolia/index.json
```

You should never commit the `.env` file, since it contains sensitive data, you can ignore it via `.gitignore`.

```
echo .env >> .gitignore
```

## Usage

This program should be ran after building site, since it need to access and upload the index file.

```
$ hugo

$ hugo-theme-bootstrap-algolia
```
