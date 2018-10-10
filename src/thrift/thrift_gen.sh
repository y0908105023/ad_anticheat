#!/bin/bash

rm -fr ./gen-go
/data/opt/thrift/bin/thrift --gen go ./filter.thrift
