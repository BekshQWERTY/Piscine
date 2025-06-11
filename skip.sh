#!/bin/bash
ls -l | tail -n +2 | sed 'n;d'