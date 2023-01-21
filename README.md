[![Go Reference](https://pkg.go.dev/badge/github.com/iwdgo/testinghello.svg)](https://pkg.go.dev/github.com/iwdgo/testinghello)
[![Go Report Card](https://goreportcard.com/badge/github.com/iwdgo/testinghello)](https://goreportcard.com/report/github.com/iwdgo/testinghello)
[![codecov](https://codecov.io/gh/iwdgo/testinghello/branch/master/graph/badge.svg)](https://codecov.io/gh/iwdgo/testinghello)

[![Build Status](https://app.travis-ci.com/iwdgo/testinghello.svg?branch=master)](https://app.travis-ci.com/iwdgo/testinghello)
[![Build Status](https://api.cirrus-ci.com/github/iwdgo/testinghello.svg)](https://cirrus-ci.com/github/iwdgo/testinghello)
[![Build status](https://ci.appveyor.com/api/projects/status/r9m4u1ew6419ikbs?svg=true)](https://ci.appveyor.com/project/iwdgo/testinghello)
[![Go](https://github.com/iwdgo/testinghello/actions/workflows/go.yml/badge.svg)](https://github.com/iwdgo/testinghello/actions/workflows/go.yml)

# Testing Hello!

Test of hello app can be done offline (server is not running) or online.
If online, it can be deployed locally (no network) or on a server.

Tests are performed:
- offline: by calling the handler directly. No server is started.  
- online without network: by issuing an http request and verifying the response of the started website.
- online with network (app is deployed): by issuing an http request and verifying the response of the deployed website.  

Use as standalone module:
- To start the site locally, use `src>go run .`

- use the standard set up (`app.yaml`) to deploy on Google Cloud which requires to create an account.  
    `src>gcloud app deploy app.yaml`

### v1.1.0

All previous solutions have been removed including comments as their use on GCP is deprecated
Previous release is tagged but requires ad hoc set up.

Further, `dev_appserver.py` does not provide support beyond go 1.11 and its use is removed.

All information regarding Google Cloud are removed as default runtime is several cycles after Go 1.11.
Repository is repurposed to test of a simple website.

### v1.0.0 Optional use of modules in various configuration.

Since `go 1.11` is available on GCP, the `app.yaml` is elementary.

## Good to know

Coverage could be improved but the required complexity is beyond the scope of this repository.

Tagging remains on v1 as the purpose of the repository is to follow the evolution of Google Cloud and go language 
on its topic.
