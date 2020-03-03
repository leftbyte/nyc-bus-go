#!/bin/bash

source cbt.rc

# QUERY="lookupVehicleInGivenHour"
# QUERY="scanBusLineInGivenHour"
# QUERY="scanEntireBusLine"
# QUERY="scanManhattanBusesInGivenHour"
# QUERY=filterBusesGoingEast
QUERY=filterBusesGoingWest

go run codelab.go -project $GOOGLE_CLOUD_PROJECT -instance $INSTANCE_ID -table $TABLE_ID -query $QUERY
