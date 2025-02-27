---
title: how minecraft host to LAN actually works 
date: 2025-02-27 08:02:06 -0500
layout: plain
---
the first time i encountered the term "LAN", i was in fourth grade trying to open up my minecraft world to my friends sitting at the computers next to me.

now, ~10 years later, as i am preparing to teach a recitation today on local area networks and ethernet, i'm revisiting how "opening your minecraft server to LAN" actually worked. and dang, it's cool!

opening to LAN does not require internet, but it _does_ require that everyone you want to play with is connected to the same router via wifi or ethernet. 

it's cool that i've forgotten all of this, but to start a LAN world, you can specify a port number and share this, alongside your PRIVATE ip address, with your friends. they can then direct connect to your server that's open to LAN using your `<private_ip_address>:<port>`.

i believe that minecraft then sends packets using UDP to whoever ends up connecting to your LAN world. if you have a bad computer, opening your world to LAN will be laggy for everyone involved, since you're effectively just hosting a server for the people in your area. 

there is something very magical about realizing that i've been interacting with all of these concepts for _years_ without knowing it.