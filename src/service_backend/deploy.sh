#!/bin/sh

gcloud app deploy --quiet --no-promote app.yaml --project=GCP_PROJECT --version=main-function
