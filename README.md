go-and-go
=========

Go link service implmeented in Go lang.

How it (should) work
====================
1. A user creates "folders" (e.g. "foo").
2. The user (owner) adds URL aliases to the folders (e.g. "foo/g" -> "http://google.com")
3. Any user registers "go" (or any word) as "Search Engine" (pointing to "http://thisservice.example.com/foo") in Chrome, and opens http://google.com by typing "go g" in the "omnibox".

Alias naming convention
-----------------------
* URL Alias must consist of [0-9A-Z-_\.]
* It can take one optional parameter

Features in beta
----------------
1. A user adds URL aliases under ".beta" namespace
2. A user resolves an alias to a URL


Endpoints
=========
GET /<user>
-> 200 returns a list of pairs of alias and URL, or
-> 404 if the name doesn't exist

GET /<user>/<name>
-> 302 with Location with the registered URL, or
-> 404 if the alias not found, or the name doesn't exist

GET /.show/<user>/<name>
-> 200 show info about <user>/<name>

POST /<user>
{ 'name': <name>, 'url': <url> }
-> 302 to '/.show/<user>/<name>'



(maybe) Useful links
====================
* https://developers.google.com/identity/protocols/OpenIDConnect
* https://developers.google.com/identity/sign-in/web/sign-in
* https://developer.github.com/v3/oauth/
* https://developer.chrome.com/extensions/tut_oauth
