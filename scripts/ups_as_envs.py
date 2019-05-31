#!/usr/bin/env python

import os, json, sys

for ups in json.loads(os.environ.get("VCAP_SERVICES", "{}")).get("user-provided", []):
  for k, v in ups.get("credentials", {}).items():
    sys.stdout.write('export %s="$(printf -- %s)"\n' % (k, repr(str(v))))
