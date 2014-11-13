Maze
======


Purpose
-------

Wish to build an peer-to-peer platform support everything possible, 
IM, social network, game, or anything else.

The platform will do the routing, tell the user where (IP address and
port) the peer he wants to connects are. Based on the address, peers 
can communicate with each others.

Based on this, many advanced feature may be added. Like a score system, 
MazeCoin, distributed DNS, or anything you can imagine.

Protocol
---
The protocol is still in draft, see it [here](https://github.com/lishaodong/Maze/blob/master/protocol/protocol.md)

Demo
----

A demo was developed in [MazeDemo](http://github.com/lishaodong/MazeDemo), it
 was build on Taipei-Torrent to find IP address, and then allow people to chat
 without any central server.

Future Work
-----------
+ current routing based on bittorrent, with tracker or DHT. For pure
peer-to-peer purpose, we hope DHT only. Current DHT used in bittorrent
is prosperous, but lack of extensibility. We may hope to proposed our 
own DHT protocol.

+ Build a platform which is extensible and easy to use. Export API only necessary.

+ Build some useful and fun Apps.

Developer Needed
----------------
If you are interested, please join me and
share you thoughts with me.

