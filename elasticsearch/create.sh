#!/bin/bash

curl  -XPUT 'http://localhost:9200/accdf' -d '
{
  "settings" : {
    "number_of_shards" : 3,
    "number_of_replicas" : 0
  },                                     
  "mappings": {
    "benchmarks": {
      "properties": {
        "Platforms>Platform": {
          "type": "nested",    
          "properties": {                                       
            "Id": { "type": "string" },                               
            "Label": { "type": "string" },
            "Version": { "type": "string", "index" : "not_analyzed" }
          }
        },                 
        "Requires>Require": {
          "type": "nested",
          "properties": {
            "Type": { "type": "string" },
            "Id": { "type": "string", "index" : "not_analyzed" }
          }
        }                                                                                                                                             
      }                                     
    },
    "testcases": {
      "properties": {
        "Tests>Test": {
          "type": "nested",
          "properties": {
            "Name": { "type": "string", "index" : "not_analyzed" },
            "Label": { "type": "string", "index" : "not_analyzed" },
        "TestSteps>TestStep": {
          "type": "nested",
          "properties": {
            "Id": { "type": "string", "index" : "not_analyzed" },
            "Ip": { "type": "string", "index" : "not_analyzed" },
            "Protocol": { "type": "string", "index" : "not_analyzed" },
            "Port": { "type": "string", "index" : "not_analyzed" },
            "Description": { "type": "string", "index" : "not_analyzed" }
          }
          }
          }
        }
        } 
  }
  }
}'
