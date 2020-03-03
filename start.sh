#!/bin/bash

### Start the GCLOUD BigTable Instance.
### Don't forget to `stop` the instance!

source cbt.rc

echo "Creating BigTable instance"
gcloud bigtable instances create $INSTANCE_ID \
   --cluster=$CLUSTER_ID \
   --cluster-zone=$CLUSTER_ZONE \
   --cluster-num-nodes=$CLUSTER_NUM_NODES \
   --display-name=$INSTANCE_ID

echo project = $GOOGLE_CLOUD_PROJECT > ~/.cbtrc
echo instance = $INSTANCE_ID >> ~/.cbtrc

echo "Creating BigTable table"
cbt createtable $TABLE_ID

echo "Creating BigTable column-family"
cbt createfamily $TABLE_ID cf
