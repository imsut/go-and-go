go-links
========

URL alias service

How it (should) work
====================
1. A user creates "folders" (e.g. "foo").
2. The user (owner) adds URL aliases to the folders (e.g. "foo/g" -> "http://google.com")
3. Any user registers "go" (or any word) as "Search Engine" (pointing to "http://thisservice.example.com/foo") in Chrome, and opens http://google.com by typing "go g" in the "omnibox".

Alias naming convention
-----------------------
* URL Alias must consist of [0-9A-Z-_\.]
* It can take one optional parameter
