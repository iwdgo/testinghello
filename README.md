[![Go Report Card](https://goreportcard.com/badge/github.com/iwdgo/htmlutils)](https://goreportcard.com/report/github.com/iwdgo/htmlutils)

[![Build Status](https://travis-ci.com/iWdGo/testinghello.svg?branch=master)](https://travis-ci.com/iWdGo/testinghello)
[![Build Status](https://api.cirrus-ci.com/github/iWdGo/testinghello.svg)](https://cirrus-ci.com/github/iWdGo/testinghello)
[![Build status](https://ci.appveyor.com/api/projects/status/r9m4u1ew6419ikbs?svg=true)](https://ci.appveyor.com/project/iWdGo/testinghello)
![Build status](https://github.com/iwdgo/testinghello/workflows/Go/badge.svg)

# Testing Hello!

Test of hello app can be done offline (server is not running) or online.
If online, it can be deployed locally (no network) or on a server.

Test is executed:
- offline: by calling the handler directly.

- online without network by issuing an http request and verifying the response of the web site.
-- To start the site locally, use `src>go run .`

- online with network (app is deployed) by issuing an http request and verifying the response of the web site.
-- use the standard set up (`app.yaml`) to deploy on Google Cloud using free quotas
    Set up of the account is required first.
    `src>gcloud app deploy app.yaml`

### Difference with v1.0.0

All previous solutions have been removed including comments as their use on GCP is deprecated
Previous release is tagged but requires ad hoc set up.

Further, `dev_appserver.py` does not provide support beyond go1.11 and its use is removed.

#### v1.0.0 Optional use of modules in various configuration.

Since `Go 1.11` is available on GCP, the `app.yaml` is very simplified.
    `src/main>gcloud app deploy .`

## Good to know

- Coverage is below standard as the 4 startup lines of main() cannot be easily tested and are 50% of the code.
The required complexities to test are outside the scope of this repository.
- Free resources are documented [here](https://cloud.google.com/free/docs/gcp-free-tier).
- Although using Go modules, the set up is confusing as reported by the build log:
`GOROOT=/usr/local/go/ GOPATH=/go GO111MODULE=on GOCACHE=/tmp/cache090267111 GOPATH=/go`
- To set `GO111MODULE=off` in GCP requires more access than usual. 
- [FAQ](https://github.com/golang/go/wiki/Modules#why-does-go-mod-init-give-the-error-cannot-determine-module-path-for-source-directory) of go modules programming reports this issue.
- After the creation of your project, no region is assigned. The first assignment is irreversible and must be
in a free quota zone:

```helloGomod>gcloud app deploy .
 You are creating an app for project [testinghello-in-the-cloud].
 WARNING: Creating an App Engine application for a project is irreversible and the region
 cannot be changed. More information about regions is at
 <https://cloud.google.com/appengine/docs/locations>.
 
 Please choose the region where you want your App Engine application
 located:
 
  [1] asia-east2    (supports standard and flexible)
  [2] asia-northeast1 (supports standard and flexible)
  [3] asia-south1   (supports standard and flexible)
  [4] australia-southeast1 (supports standard and flexible)
  [5] europe-west   (supports standard and flexible)
  [6] europe-west2  (supports standard and flexible)
  [7] europe-west3  (supports standard and flexible)
  [8] northamerica-northeast1 (supports standard and flexible)
  [9] southamerica-east1 (supports standard and flexible)
  [10] us-central    (supports standard and flexible)
  [11] us-east1      (supports standard and flexible)
  [12] us-east4      (supports standard and flexible)
  [13] us-west2      (supports standard and flexible)
  [14] cancel
```


Releases below are for illustration purposes. Using the latest releases is recommended.

```

go version go1.13.3

>gcloud components list

Your current Cloud SDK version is: 268.0.0
The latest available version is: 268.0.0

| Installed     | App Engine Go Extensions                             | app-engine-go            |  4.8 MiB |
| Installed     | BigQuery Command Line Tool                           | bq                       |  < 1 MiB |
| Installed     | Cloud Datastore Emulator                             | cloud-datastore-emulator | 18.4 MiB |
| Installed     | Cloud SDK Core Libraries                             | core                     | 12.1 MiB |
| Installed     | Cloud Storage Command Line Tool                      | gsutil                   |  3.6 MiB |
| Installed     | Google Container Registry's Docker credential helper | docker-credential-gcr    |  1.8 MiB |
| Installed     | gcloud app Python Extensions                         | app-engine-python        |  6.0 MiB |
| Installed     | kubectl                                              | kubectl                  |  < 1 MiB |
```
