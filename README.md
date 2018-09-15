**Testing Hello! locally**

Test of hello app can be done offline (site is not running) or online.
If online, it can be deployed locally in three different ways.

Offline: by calling the handler directly
Online: by issuing an http request and verifying the response.

To start the site locally, you can:
- run main() and `ListenAndServe()`
    
    `src>go run ./helloFlex/helloFlex.go`
- use the standard set up (`appStd.yaml`)
    
    `src/helloStd>dev_appserver.py appStd.yaml`
- use the flex set up (`appFlex.yaml`)
    
    `src/helloFlex>dev_appserver.py appFlex.yaml`

Notice the change of default version depending on the method used.
The handler is the same as it only prints the chosen sentence.
It cannot be imported from a common package as ../ is not supported by dev_appserver.
Other differences in set up between Std and Flex deployment are commented.

The repo is using a recent gcloud SDK with default components installed:

```
| Cloud SDK Core Libraries               | core                     |   8.5 MiB |
| App Engine Go Extensions               | app-engine-go            | 154.5 MiB |
| BigQuery Command Line Tool             | bq                       |   < 1 MiB |
| Cloud Datastore Emulator               | cloud-datastore-emulator |  17.7 MiB |
| Cloud Storage Command Line Tool        | gsutil                   |   3.6 MiB |
| gcloud app Python Extensions           | app-engine-python        |   6.2 MiB |
| kubectl                                | kubectl                  |   < 1 MiB |
```
