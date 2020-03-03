# Summary
A Go rewrite of the Java based codelab sample of the nyc-bus example from
https://github.com/googlecodelabs/cbt-intro-java.

The nyc bus app is an example that demonstrates how to use Google's
Cloud BigTable as a data backend for efficiently storing and
retrieving bus location data.

# Quick  Setup

For full details, please visit Google's Introduction to Cloud Bigtable:
https://codelabs.developers.google.com/codelabs/cloud-bigtable-intro-java

This is a summary of what I do to instantiate the
infrastructure. Having your gcp account and installing `gcloud` and
`cbt` are necessary prerequisites.

First modify cbt.rc and set GOOGLE_CLOUD_PROJECT according to your
configuration. Next set your GOOGLE_APPLICATION_CREDENTIALS.

E.g. `export GOOGLE_APPLICATION_CREDENTIALS="/Users/leftbyte/.gcp/nyc-bus-123.json"`

Then run the following:

$ `bash start.sh`

$ `bash import-data.sh`

Now you should be able to modify and run the sample:

$ `bash run-codelab.sh`

When you're done, don't forget to shutdown your services so you don't
run up the bill!

$ `bash stop.sh`
