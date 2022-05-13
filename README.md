# flashy - WIP

An app to help you learn English better!

Type in your words and sentences into these flashcards, then try to memorize what you just wrote. You will be surprised at how efficient this is.

## UI/UX

<img src="/images/output/video1.gif" width="50%" height="50%"/>

**_The app flow is simple and intuitive:_**

_*I'm not a designer so somehow the cards look like a floppy disk*_ :blush:

- Sign In
- Choose how many words you want to learn
- Learn them
- Then go back to home to see your new words appears in your history

**_Features:_**

- **Auto Login:**

Log in once and never have to worry about it ever again!

- **History:**

Keeps a record of every words you have learned.

To be added: your wrong guess.

- **Versus mode - WIP:**

Allow players to compete head to head thanks to gRPC's amazing bidirectional streams. Currently the Go server is completed. Now we only need to add support this in the gRPC Java module.

- **Rankings - WIP:**

Knowing how many words your friends have learned to date.

### TODOs:

- Deploy somewhere cheap (preferably locally on a Raspberry-Pi)
- Deploy to Playstore
- Friends, rankings
- Client versus mode (gRPC bidirectional streams in Java)
- UI for friends, leaderboards, versus.
