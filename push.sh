#!/bin/bash

theSpace="$(cf target)";

echo "Details of the space you are in: $theSpace";
echo ""

if [[ $theSpace == *"-prod"* ]]
then
    echo "Pushing to Production Space";
    echo ""
    cf push
elif [[ $theSpace == *"-staging"* ]]
then
    echo "Pushing to Staging Space";
    echo ""
    cf push -n px-application-name-staging -f ./manifest_staging.yml
elif [[ $theSpace == *"-qa"* ]]
then
    echo "Pushing to QA Space";
    echo ""
    cf push -n px-application-name-qa -f ./manifest_qa.yml
elif [[ $theSpace == *"-dev"* ]]
then
    echo "Pushing to Development Space";
    echo ""
    cf push -n px-application-name-dev -f ./manifest_dev.yml
fi
