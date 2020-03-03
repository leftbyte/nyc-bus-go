#!/bin/bash

### Import data into BigTable

source cbt.rc

echo "Enabling dataflow on google API"
gcloud services enable dataflow.googleapis.com

echo "Starting import of data"
NUM_WORKERS=$(expr 3 \* $CLUSTER_NUM_NODES)
gcloud beta dataflow jobs run import-bus-data-$(date +%s) \
--gcs-location gs://dataflow-templates/latest/GCS_SequenceFile_to_Cloud_Bigtable \
--num-workers=$NUM_WORKERS --max-workers=$NUM_WORKERS \
--parameters bigtableProject=$GOOGLE_CLOUD_PROJECT,bigtableInstanceId=$INSTANCE_ID,bigtableTableId=$TABLE_ID,sourcePattern=gs://cloud-bigtable-public-datasets/bus-data/*
