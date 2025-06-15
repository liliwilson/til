---
title: i finally read the original google paper! 
date: 2025-06-15 16:56:36 -0400
layout: plain
---
somehow, even after heavily focusing on computer systems in undergrad (and fixating on _networking_ and _distributed systems_ in particular), i am only now reading the original google paper (by sergey brin and larry page, 1998). but i'm so glad i did—it was a really fun paper to read, and surprised me more than i thought that it would. 


some thoughts i had while reading this paper:

### it's hard to imagine a world where google was just an "academic prototype."
seeing "try out our search engine at http://google.stanford.edu" was pretty wild. 

the paper was also quite critical of companies developing search engines: "up until now most search engine development has gone on at companies with little publication of technical details. this causes search engine technology to remain largely a black art and to be advertising oriented." 

i wonder if this mentality is still a part of conversations at google now. as a user, i definitely feel a massive influence of advertising on my experience of the search engine. as a computer systems person, i also feel the increase in black box-iness of google's systems (for example, the transition away from the publicly-described GFS to the much more secretive colossus). i would be curious to hear original googlers' take on this perceived change, and whether or not the internal mentality has been consciously shifting as well (and if so, why!).

sergey and larry's commitment to creating a public description of their search engine was really cool, and although more recent google advances feel much more Under Wraps, it's neat that we still have this original version of the system to engage with. their vision of google as a research tool was also quite intriguing, and they seemed to emphasize this quite a bit when they justified google's storage of entire webpages, not just the parsed data for the search indices.

### this paper embodies what i love about computer systems, and the type of systems i hope to build.
okay, saying "i want to build a system like google" is _really_ not radical, i know. but what i loved about this paper, and the initial visions for the system, and my desire to be a part of systems like it in the future, goes a bit deeper than wanting to make a wildly successful and globally impactful tool.

it's incredibly beautiful to me how this paper outlines what feels like a single key innovation that makes google unique: using the hyperlinks in documents to assess the "quality" of a search results. this single, core idea is paired with a ton of nifty supporting optimizations that power it, but at it's core, it's quite an accessible and understandable concept. the system's potential does not lie in complex, opaque ideas, but instead, in _simplicity_ that is done _elegantly_. the complexity lies in the implementation details that make it performant and scalable. these are the kinds of systems i want to build (my most recent experience with something like this was building a file distribution system in my networking class last fall!)

and speaking of the implementation details, a lot of the optimizations talked about in this paper are so _neat_: the encoding schemes for hitlists (2 bytes per hit!), tradeoffs considered when ordering docIDs in doclists, handling of the DNS bottleneck for crawler performance...

one last thing that i love — i love a computer system that takes inspiration from the authors' experience of the rest of the world. linking pagerank to the way academic citations work was neat! and explaining pagerank with the "random surfer" analogy was both compelling and fun.

### they talk about the people involved.
my favorite systems papers talk about people, as well as computers. this paper does a great job grounding the technical in the real people that want to use search engines, which i enjoyed.

including information about the social issues involved in web crawling was also awesome. it made the paper feel more human. and it was something i hadn't considered before (especially the part about an excited webmaster thinking someone had been eagerly engaging with every link on their site, only to find out it was a bot...)

### this paper is funny.
i love academic paper humor. my favorite line from the paper comes from their evaluation of google's search results for the query "bill cilnton":  "there are no results about a bill other than clinton or about a clinton other than bill." a close second is the inclusion of the stat that "most search engines cannot even locate themselves."
