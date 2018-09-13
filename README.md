***Testing Hello World! locally***

Testing sending hello to the world will work can be achieved in various ways.

Offline:
 - calling the handler directly

Online:
The page is reached using and http request.
Running the site can be done:
- running main and ListenAndServe()
    go run hello...
- using the standard set up (appStd.yaml)
    dev_appserver.py appStd.yaml
- using the flex set up (appFlex.yaml)
    dev_appserver.py appStd.yaml

Notice the change of default version depending on the method used.