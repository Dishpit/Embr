---
sidebar_position: 7
---

# Hashes

As of v0.9.0, Omega has functional hashes.

## Using Hashes

Currently, hashes can only be used through live interpretation, and are unable to be stored as variables due to a lack of a hash data type.

### Syntax

The basic syntax for using hashes in Omega is as follows:

```omega
{"name": "Omega", "based": true, 420: 69, "Zaphod": "Beeblebrox" }["Zaphod"]; // this returns: "Beeblebrox"
```
