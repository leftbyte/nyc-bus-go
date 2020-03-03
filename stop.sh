#!/bin/bash

### Stop the GCLOUD BigTable Instance.

source cbt.rc

echo "Deleting bigtable $INSTANCE_ID"
gcloud bigtable instances delete $INSTANCE_ID
