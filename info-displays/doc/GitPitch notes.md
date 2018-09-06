## Overview

GitPitch is GitHub based slide presentation, by GitHub.

See the [60 second overview](https://github.com/gitpitch/in-60-seconds) to get started - it really is that easy.

### Multiple Pitches in a single repo

You can store different pitches in different branches (default), OR different directories:

* Different branches: `https://gitpitch.com/$user/$repo/$branch`
* Different directories: `https://gitpitch.com/$user/$repo/$branch?p=$directory`


View this pitch at [https://gitpitch.com/stevetarver/rft-presentations/master?p=svgo](https://gitpitch.com/stevetarver/rft-presentations/master?p=svgo)

### GitPitch links

* [Features](https://gitpitch.com/features)
* [Docs](https://gitpitch.com/docs)
* [Template Gallery](https://gitpitch.com/templates)
* [Kitchen Sink](https://gitpitch.com/gitpitch/kitchen-sink#/)

## Tips

### Rapid development

To shorten the write-push-view loop:

1. Create a skeleton one-page presentation and git-push it.
1. Open your skeleton one-page presentation on gitpitch.com.
1. Click the burger-menu and then "Offline Version" to download the self-contained presentation bundle.
1. Follow instructions from [here](https://gitpitch.com/docs/foundation-features/offline/) to run the presentation. E.g.
    1. Copy the `PITCHME.zip` file to your project directory and expand it, creating a `PITCHME` directory.
    1. `cd PITCHME`
    1. `python3 -m http.server`
    1. Edit the pitch at `assets/md/PITCHME.md`
    1. Reload the pitch after edit to view the results.
1. When edits are complete, merge the above into your main PITCHME.md

### Domain markdown

https://gitpitch.com/docs/markdown-features/shortcuts/

### Background images

The slide delimiter (`---`, `+++`) marks the beginning of a slide. Add an image to that and it becomes the background image.

```
---?image=svgo/pitch/images/gopherhat.jpg
```
Not the similarity to a url query string.

The path is relative to the repo root - don't forget the sub-directory if you are using a single repository for multiple presentations.

* SVG images require purchasing the 'pro' package.


### Add an image to a slide

### Formatting

