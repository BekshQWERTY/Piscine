#!/bin/bash
grep " 170 " superhero | awk '{print $2 "\n" $3 "\n" $4}'