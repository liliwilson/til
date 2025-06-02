---
title: CAPTCHA is an acronym? 
date: 2025-06-02 14:38:59 -0400
layout: plain
---
and quite a funny one at that.

CAPTCHA stands for "Completely Automated Public Turing Test to tell Computers and Humans Apart." thank you to the may 31, 2025 nyt crossword for drawing this to my attention (:

while looking up this acronym, i stumbled across some more interesting CAPTCHA-related information, which i shall share here.

although CAPTCHAs were originally developed to distinguish between humans and computers for security, they have since become very useful for AI. problems that make for good CAPTCHAs are usually the ones that current AI _cannot_ solve. using "hard AI" problems for CAPTCHAs then serves two purposes: 1) providing security in the event that the problem goes unsolved, and 2) helping us continue to make progress on these "hard AI" problems either by having humans do some of the labor or by generating good training data for models. 

but what problems make good CAPTCHAs? turns out, there's a lot of interesting work that has been done here, too. 

- reCAPTCHA (a form of CAPTCHA that emerged out of CMU but was soon bought by Google) has been used to digitize New York Times articles and Google Books. users are shown two words, one of which needs to be digitized (likely from an old scan of a book or newspaper).
- all of the CAPTCHAs that involve identifying traffic lights, fire hydrants, or motorcycles, are likely used to train self-driving vehicles.
- i saw a CAPTCHA earlier today that asked me to identify "things that are naturally biodegradable". did i help train a trash bot??
- [this paper](https://link.springer.com/content/pdf/10.1007/3-540-39200-9_18.pdf) expands on the idea of using "hard AI" problems for CAPTCHAs, and defines a CAPTCHA family called MATCHAs. i love good names.

and of course, there are multiple [xkcd](https://explainxkcd.com/wiki/index.php/1897:_Self_Driving) [comics](https://explainxkcd.com/wiki/index.php/2604:_Frankenstein_Captcha) about this.

CAPTCHAs are becoming harder and harder to develop, though, and we are already past the point where most humans are worse at the challenges than AI. google rolled out No CAPTCHA reCAPTCHA in 2014 to help with this, and reCAPTCHA v3 in 2018, which both aim to distinguish between humans and bots without requiring challenges to be filled out. v3 is designed to be "invisible" to most users, which uses other data about a user's internet activity to try to determine whether or not they're a bot. 

i am curious to see whether CAPTCHAs continue to be used, and whether we'll hit a point where it's truly not possible to distinguish between humans and bots online in the ways that we have been doing. are we already there? ("proof of personhood" online is becoming a _massive_ question, especially since 2022 with the release of major text generating LLMs...)
