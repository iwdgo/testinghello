#Testing Hello! locally

Test of hello app can be done offline (site is not running) or online.
If online, it can be deployed locally in three different ways.

Offline: by calling the handler directly
Online: by issuing an http request and verifying the response.

### Go 1.11

Since `Go 1.11` is available on GCP, the `app.yaml` is very simplified.
    `src/main>gcloud app deploy .`

### In general

To start the site locally, you can:
- run main() and `ListenAndServe()`
    
    `src>go run ./helloFlex/helloFlex.go`
- use the standard set up (`appStd.yaml`)
    
    `src>dev_appserver.py ./helloStd/appStd.yaml`
- use the flex set up (`appFlex.yaml`)
    
    `src>dev_appserver.py ./helloFlex/appFlex.yaml`
- use the standard set up (`appStd.yaml`) to deploy on Google Cloud using free quotas
    Set up of the account is required first.
    `src/helloStd>gcloud app deploy appStd.yaml`

Default version of Go depends on the deploy method.
The handler is the same as it only prints the chosen sentence.
It cannot be imported from a package as ../ is not supported by dev_appserver.
Other differences in set up between Std and Flex deployments are commented.

Check release notes for detail if you use a version before 1.11

The repo requires a recent gcloud SDK with default components installed:

```
go version go1.11.4 windows/amd64
| Installed     | App Engine Go Extensions                             | app-engine-go            |  56.9 MiB |
| Installed     | BigQuery Command Line Tool                           | bq                       |   < 1 MiB |
| Installed     | Cloud Datastore Emulator                             | cloud-datastore-emulator |  17.7 MiB |
| Installed     | Cloud SDK Core Libraries                             | core                     |   9.1 MiB |
| Installed     | Cloud Storage Command Line Tool                      | gsutil                   |   3.5 MiB |
| Installed     | Google Container Registry's Docker credential helper | docker-credential-gcr    |   1.8 MiB |
| Installed     | gcloud app Python Extensions                         | app-engine-python        |   6.2 MiB |
| Installed     | kubectl                                              | kubectl                  |   < 1 MiB |
```
