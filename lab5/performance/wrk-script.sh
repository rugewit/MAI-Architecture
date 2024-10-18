#!/bin/bash
wrk -t12 -c400 -d30s http://0.0.0.0:8081/users/665662d9ea106dc5807716bc
