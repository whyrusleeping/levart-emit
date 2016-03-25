# levart-emit
> (i know, creative, right?)

This is a simple tool for moving content between ipfs nodes in a way that
doesnt care about versions. Got stuff on 0.4.0 you want to move quickly to
0.3.11 (without mucking with migrations)? Use levart-emit. Want to push some
stuff from your 0.3.11 node up to an 0.4.0 node? Uh... well you can't quite use
levart-emit for that, seeing as the name is time travel backwards, and moving
from 0.3.11 to 0.4.0 is moving forwards in time, which isnt super spectacular.
So I didn't make it do that. (if you really want to, you can change the hosts
in the code and recompile it)

## Installation
```bash
$ go get github.com/whyrusleeping/levart-emit
```


## Usage
I'll probably work more on this if anyone actually uses it, but for now the usage goes something like this:

1. run an ipfs node (with api on port 5001, i'm lazy I hardcoded that)
2. find content on the 0.4.0 network you want on your node.
3. `levart-emit <HASH OR PATH>`

## Troubleshooting

> "Its going really slow!"

Time travel is hard. (and we don't do anything in parallel).

### License
MIT
