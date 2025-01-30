# til
a little cli tool for posting to my own rss feed!

very much a work in progress, and very much janky at the moment.

inspired by [nom](https://github.com/guyfedwards/nom?tab=readme-ov-file) and various TIL blogs that i have really been enjoying lately (in particular, [this one](https://jvns.ca/blog/2024/11/09/new-microblog/) from julia evans!)


## how does it work?

til is built in go, using [cobra](https://github.com/spf13/cobra) for parsing command line input and [bubbletea](https://www.google.com/search?q=bubbletea&sourceid=chrome&ie=UTF-8) for rendering the tui.

using til is very simple!

type `til` into your cli to get this nice little editing screen:


<img width="552" alt="til command" src="https://github.com/user-attachments/assets/8f1943c0-6491-4159-9609-608634bcb1d9" />
<img width="552" alt="Screenshot 2025-01-29 at 8 52 23 PM" src="https://github.com/user-attachments/assets/bb6afff5-6b99-4e97-a02c-7a686b0e011a" />

<img width="552" alt="til edit" src="https://github.com/user-attachments/assets/4a3c524c-7aba-4848-8965-0394c63546bf" />

then, `til compile` will generate an rss feed and `til publish` will send it over to this github repo. sorry, i have not made this usable for other people yet :p

## anything else cool?
yep! i have confirmed that it works with `nom`, which is also neat.

<img width="552" alt="filtering nom" src="https://github.com/user-attachments/assets/d98f3cfb-c5e3-4f8f-b2e2-627c1aae044e" />
<img width="552" alt="reading my thing" src="https://github.com/user-attachments/assets/edfe4bbc-698f-4e9d-851f-d4b64c619c61" />

i am indeed subscribed to my own rss feed.


