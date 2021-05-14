# Git GOod Discord
Group 1 - The Three Gophers

- IP: 10.212.136.172
- Port: 80
- NTNU VPN required

### Group members
- Simen Bai (simenbai@stud.ntnu.no)
- Ruben Christoffer Hegland-Antonsen (rubench@stud.ntnu.no)
- Even Bryhn Bøe (evenbbo@stud.ntnu.no)


## Setup

The system is not intended to be run locally although it is possible.

1. Clone repository with `git clone https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2021-workspace/simen_bai/git-good-discord-group-1.git`
2. Add firebase and discord credentials from the example files
3. Update IP of service.json
4. Firestore needs a composite index on **channel_id** and **id** for the **subscriptions** command.
   This can either be added manually, or you can try the **subscriptions** command which will crash the system.
   The error message will contain a link to generate the index.
5. Run docker `sudo docker compose up -d --build`


## Usage

There is a [landing page](http://10.212.136.172) with the same basic user information as in this section.

### Pre use requirements

For some commands to work the user executing them needs the **Admin** role (case sensitive).
It is therefore strongly recommended for a channel to have an **Admin** role.

The system is run on OpenStack and a VPN is required to reach the IP.

### Setup

- Authorize the bot for the server through [the discord website](https://discord.com/oauth2/authorize?client_id=830135256514297936&scope=bot)
- Add **Admin role**

### Default settings

There are some default setting for each channel that can be changed later on.

- The default language is **english**
- The default command prefix is **!**

### Discord commands

Assuming default command prefix (**!**)

`<item>` is a placeholder where **item** indicates what should go in the placeholder

#### General commands

- `!help` Lists commands
- `!subscibe <gitlab_instance>/<repo_id>/<gitlab_username> <type1,type2>` to subscribe to events
    - This requires that an access token has been added for the project
    - **type** is the type of event given as list e.g. issues,merge_requests or issues
    - e.g. `!subscribe git.gvk.idi.ntnu.no/1965/evenbbo1 issues,merge_requests`
- `!unsubscribe <gitlab_instance>/<repo_id>/<gitlab_username>` to unsubscribe from all events
    - e.g. `!unsubscribe git.gvk.idi.ntnu.no/1965/evenbbo1`
- `!subscriptions` to get subscriptions of channel 

#### Admin commands

Commands requiring user to have the role **Admin**

- `!reload` to reload language packs. This is not usually needed but allows for changes in language packs without having to reboot the system.
- `!language <language>` to set language of channel
    - e.g. `!language english`
- `!set prefix <prefix>` to set prefix of the channel
    - e.g. `!set prefix #`
    - Note: This command always requires the preifx **!**. This is in case the prefix is forgotten
- `!access_token <instance>/<repo_id>/<token>` set access token of a project
    - specific to each channel

## Known bugs

- Nothing that we know of


## Further improvements

- Gitlab has more events such as push, comments, pipeline.
  This could be added, although it would mostly be similar to merge requests
  and issues and would therefore not demonstrate any greater competence.
- Caching. To reduce the amount of requests sent to the database as we are 
  limited to 50k reads per day (not an issue in the scope of this task but it could be in a large scale deployment).
  
  The settings for each channel is read for each message sent. Settings are one thing that 
  we only really needed to read once on startup and then the only interaction with firebase would be writing.

- Testing. We could add more tests. There are enough tests demonstrate competency 
  in testing and mocking, but we have not added tests to everything.

## Limitations

As of now it only allows for connection to the **git.gvk.idi.ntnu.no** gitlab instance. 
This is not a limitation of the program but rather of the deployment method. 
Since we are using OpenStack other instances can not reach the IP and therefore not post anything. 
Switching to a different IaaS such as GCP would fix this 
(could also be a PaaS solution as long as it has support for docker as our deployment method uses it).


---

# Report

## Project plan

The idea was to create a discord bot that would allow users to subscribe to gitlab events.
Users can choose a gitlab user in a repository to subscribe to 
and when that user is mentioned in an event; a message is sent to discord to notify the user 
with the discord mention functionality (using **@username**).


## What went wrong

Time. We had a project in Computer Vision due the 10th, 
and an oral exam in Computer Vision the 18th which stole some time and attention.


## What was hard

- Dealing with cyclic imports.
- Merge conflicts


## Learning outcome

- We learned about the quirks of Go. Especially in regard to the cyclic imports
- Working a lot on the same files resulted in some good practice with merge conflicts. 
- Experience in how to work in a team on a larger project
- Good experience in reading API documentation. 
  We chose well documented APIs (discord and gitlab) 
  but reading through and applying the necessary changes for our use 
  is still a valuable skill to have even if we were spared some headaches from bad documentation.


## Work hours

We forgot to count hours although we have been working some hours every day for the last few weeks.

An educated guess would be **170 hours** in total.
