
---
title: "Details of the new TPF blog site"
author: Henry Van Styn
type: post
date: 2019-10-03 00:00:00 +0000 UTC
url: "/post/new_tpf_blog_details"
categories:
 - Perl Foundation
 - Sign in

---

Late this evening we flipped the switch and The Perl Foundation News website (http://news.perlfoundation.org) no longer lives at its long time, MoveableType-based home, and is now a brand new system based on [Rapi::Blog](https://metacpan.org/pod/Rapi::Blog), which is a relatively new blogging platform, based on [RapidApp](http://rapi.io), written in Perl, open-source and distributed on CPAN.

Save for the import of users and posts from the old system, the new site has been built from the ground up, and so everything from a functionality standpoint is completely new, which means everything works a bit differently now. 

Hopefully most of the new functionality is self-explanatory, but I'm going to outline the major points of what you need to know in order to understand and get around the new system as painlessly as possible.


### About the Rapi::Blog platform

If you're interested in learning more about Rapi::Blog and what it is all about, I gave a 50-minute talk on it at the 2017 Perl Conference, which was the day of its very first public release. You can find the talk video, abstract and slides here:

  * http://www.rapidapp.info/demos/tpc2017

Obviously a lot of improvements have been made since then, but the main points and the goals of the project remain the same. During the talk I did lots of live demos and walkthroughs of interfaces, so that might prove helpful in seeing how stuff works by example.

# New design (for now)

If you're an ordinary visitor who just comes here to read the latest Perl Foundation news, the only difference you'll notice is the new design, and a slightly different layout. Currently this is nothing fancy, as it is more or less just the default Bootstrap-based design that ships with Rapi::Blog. 

We still plan to create a proper, branded design, but for now the current design is just meant to be functional with no-frills. One of the key features of Rapi::Blog is that it is easy to swap layouts/designs (called "scaffolds") so it is no problem to do the redesign as a separate step and largely independently. 

If you're interested in helping out with this effort, read on to the Contributing section.


# Logging into your account

If you had an account on the old system, it has already been migrated to the new system. However, you will need to activate your account and set a new password. To do this, click "Sign in" and then click "E-Mail Login" and enter either your username or your E-Mail address. The system will e-mail you a one-time link that you can follow to login without a password. 

Once you're logged in, click on your username to open your account profile page. To set your password, simply click on the Set Password field and follow the prompts to enter and confirm you password that you can use to login normally going forward.

If the system says that there is no e-mail on file for your account, an admin will need to enter it for you. The export file I received didn't have e-mail addresses for a lot of users. If this is you, don't fret, you can just e-mail me directly and I'll update your account (my contact info is at the end of the post).

# Posting articles

For users with the 'author' permission, click "New Post" to load the new post form which has the following fields:

 * **Name**: This must be unique and is used to generate the URL for the post. It is limited to sane values in order to ensure clean URLs. It won't allow you to enter spaces, special characters, or capital letters
 * **Title**: This is the human friendly title which can be whatever you want. If you leave it blank the system will just use the same value as name
 * **Categories**: Select zero or more categories for your post. This is just a taxonomy to help group and organize posts, same as many other systems
 * **Date/Time**: set the timestamp that you would like to appear in the public view of the article. Note that the system still tracks other immutable timestamps, such as ``created_ts``, ``updated_ts``, ``published_ts``
 * **Published (<img 
  src="_ra-rel-mnt_/simplecas/fetch_content/7a819a65fbfe27f1e2c079efbfd5cff0024fc25e/eye.png" 
  style="max-width:100%;"
 />)**: Set to yes or no to control whether or not the post will be publicly visible. This allows you to work on a post, coming back to it across multiple sessions, and not make it live until you're ready
 * **Image**: This is an optional image which you can set which will appear with the article summary
 * **Custom Summary**: You can fill this out to set the summary of your post to whatever you like. If you leave it blank, an auto summary will be generated, which will simply be a snippet of the first few lines/paragraph
 * **Tags**: In Rapi::Blog, tags are automatically set by parsing out all hashtag style keywords from the body of your article. For instance, if the string #RapidApp is anywhere within your article, it will automatically be tagged with "RapidApp". If later on you edit your article and remove #RapidApp, the RapidApp tag will automatically be removed

## Post body

The Post body is where you actually compose your article. The editor is very robust with many features including rich content editing, drag drop embed of images and even file attachments, as well as other bells and whistles. 

You can compose articles using markdown as well as HTML, and the editor will actually update the formatting in real-time to give you a better feel for how your content is shaping up. It is important to note that this is **NOT** a WYSIWYG which munges your original input into a garbled html mess which can't be reversed. This is still just a __text editor__ and it is still just text that is displayed, but with different text rendered with different fonts colors and sized. It is basically just a more aggressive syntax highlighter.

This is helpful, but to get an even better sense of what your marked up content will actually look like when rendered, you can click "side-by-side preview" to toggle split-screen mode which will show how your content will be rendered on the right side in real time as you type your  content on the left.

This gets very close to the actual final look of your article, but still not 100%, because in order to see EXACTLY what it will look like when viewed on the actually, public-facing website, all the surrounding CSS and styles must be factored in.

But, the system has you covered here, too. If you save your article unpublished, it loads you to the "Edit Post" page where you can see your post fully rendered via the public site, and this is *exactly* how it will look once published. This is done using an iframe (note: as this iframe preview panel is meant as a view-only resource, the system blocks any link click/navigation events. You can still click links, they just won't do anything. This is by design)

Note that as the author of the post, you will be able to see it via the public URL, even when unpublished. But, all other unprivileged users will get a 404 until you publish it.

On the Edit Post page you'll see the conspicuous **[NOT PUBLISHED]** banner. But within the iframe and rendering of the post, there will be a little red eye with a cross through it next to the title, indicating it is unpublished.

<div style="padding-left:100px;padding-bottom:30px;">
<span style="opacity:0.7;vertical-align:middle; font-size:1.3em;padding-right:8px;"><i>Unpublished indicator:</i></span>
<img style="vertical-align:middle;" 
  src="_ra-rel-mnt_/simplecas/fetch_content/e8acad32b41dafbbe6b7d4e7124e2409be8de858/red_eye.png" 
  style="max-width:100%;"
 />
</div>

## Edit Post page

As I already covered in the previous section, the Edit Post page displays the native final rendering of the post by using an iframe. You can then "Click to Edit" to be taken back to the same Post body editor with all its rich features. You can then save, view, edit, save, view, as many times as you like until you're ready to publish.

To mark a post as Published, you need to click the "Published" toggle at the top of the Edit post page, towards the right-hand side.

The Edit Post page is actually a tabbed interface with 3 tabs. We've already seen the first two and the third is _Attributes_. This tab allows you to view and edit the other datapoints tied to the post, such as name, title, categories, timestamp, and so on.

Note that certain fields will be read-only, or even hidden altogether, depending on user permissions. For instance, admins which are the highest level are all powerful, so they can do things like create posts as other users, while everyone else can only create posts as themselves. As a result, for admins, they'll have an editable "Author" field on the new post page, which is completely hidden from everyone else. However, it should still be noted that even though admins have these powers, they still can't edit fields like ``updater``, ``update_ts``, etc. So, while they can do lots of things, they can't suppress the record of them doing it.


# Admin section

There is also a backend admin area which is only accessible to authenticated users. This can be accessed by clicking the gear icon at the top-left of the public site.

This area provides access to all the same data as the public interfaces, plus a lot more, and all of it can be queried and manipulated via much, much more powerful interfaces, all powered by RapidApp.

I'm not going to take too much time to dive into this right now. But this is the area where all the administrative functions takes place, so this applies much more to admins. However, even normal users can utilize this section by being able to build complex queries to explore the data and relationships in just about any manner they can imagine.

# Public Repositories

RapidApp, Rapi::Blog and tpf-blog-live (the blog site itself) are all open-source with public Git repositories:

 * https://github.com/vanstyn/RapidApp
 * https://github.com/vanstyn/Rapi-Blog
 * https://github.com/RapidApp/tpf-blog-live

Patches and discussions are always welcome

# Running your own local version
If you'd like to fire up your own instance of the TPF blog for testing and experimenting, it is very simple. The recommended way is with docker, since this is usually fastest and most reliable. To make this even more dead easy, you can use my [dockup.sh](https://github.com/docker-rapi/dockup.sh) sugar script, which is a drop-in replacement for ``plackup``:

```
wget -O - http://rapi.io/install-dockup.sh | bash
docker pull rapi/psqi
git clone https://github.com/RapidApp/tpf-blog-live.git
cd tpf-blog-live/
dockup.sh
```

Or, you can use the more traditional setup using your local perl, 

```
cpanm Rapi::Blog
git clone https://github.com/RapidApp/tpf-blog-live.git
cd tpf-blog-live/
plackup
```

But beware this could take a long time and is less consistent as sometimes upstream CPAN modules break. This is why I developed and maintain the [rapi/psgi]( https://cloud.docker.com/u/rapi/repository/docker/rapi/psgi) Docker Hub image. It is a clean, pre-installed and pre-validated module stack so it always works and only takes seconds to download.

But, with either of the above commands, you'll have a fully working local instance (excluding content, so it will start off empty, and you'll need to create your own accounts and posts).

Note that freshly initialized databases start off with the admin account with the password 'pass' (first thing you should probably do is change that)

# Contributing

The code base for the TPF blog site itself is available here:

 * https://github.com/RapidApp/tpf-blog-live.git

If you browse its code, you'll find that there really isn't that much there. The majority of the files are in the ``scaffold/`` directory and this is the full layout and design of the site. A different scaffold, and thus an entirely different design can be activated simply by replacing the current scaffold directory for another.

The scaffold is structured as a plan/static HTML site, and like a static site, the directory paths and files correspond to the URLs served. It contains html files, css files, js files, images and so on in the same structure as a static HTML site. In fact, an actual, static HTML directory structure would function as a scaffold just fine.

But, of course, scaffolds typically serve dynamic content, and this is achieved by adding template directives based on [Template::Toolkit]( https://metacpan.org/pod/Template::Toolkit) to the html (and other text files) within the scaffold. While Template::Toolkit is the engine used by Rapi::Blog, a powerful and comprehensive API is provided in the form of custom template directives which are available to be called by any of the files within the scaffold. See the [TEMPLATE-DIRECTIVES](https://metacpan.org/pod/Rapi::Blog::Manual#TEMPLATE-DIRECTIVES) section of the Rapi::Blog manual.

As I mentioned earlier, the current site design is really just one of the built-in scaffolds that ships with Rapi::Blog. And, the purpose of this is not only to provide a working site out of the box, but also to serve as a _reference implementation_. It is a scaffold which calls most of the available template directives provided by the API, but also tries to use them in the cleanest, and most straightforward way possible, so that it is easy for a new designer/developer to see how to use the API to incorporate into a new design.

So, all that being said, contributions are welcome and encouraged, particularly regarding the design. One of the main purposes of the Rapi::Blog project was to make the process of implementing front-end designs and layouts as easy as possible, with the goal of making it so accessible, that even a designer who knows Photoshop/HTML/CSS but isn't a coder can do it.

Besides designs, any and all contributions are welcome. To submit a patch to any part of the stack, just open a pull request.

Contributions aren't limited to code patches. I'm also more than open to feature suggestions, bug reports, general comments, etc. Also, documentation contributions would be especially helpful. 


# Further reading/resources

 * [Original 2017 Rapi::Blog talk](http:// 
http://www.rapidapp.info/demos/tpc2017)
 * [Rapi::Blog Manual](https://metacpan.org/pod/Rapi::Blog::Manual)
 * [RapidApp Manual](https://metacpan.org/pod/RapidApp::Manual)
 * [RapidApp project homepage](http://rapi.io)

<br>

# Final Thoughts

I know this article ended up being quite long, so if you've read this far, kudos! As long as this is, I still only managed to scratch the surface. Rapi::Blog has been under development for around 2 years, but RapidApp goes back much further (and, in case you're curious, yes, client projects built on RapidApp is how me and the rest of my team at IntelliTree earn a living).

RapidApp has been under continuous development now for over 10 years, and it has grown to have a massive arsenal of features. This is actually both RapidApp's greatest strength as well as its greatest weakness, as it does so much it can be daunting for new users. So most people use RapidApp as a glorified CRUD tool, but it is far larger than that.

Anyway, I'm very excited to have the TPF blog now running the new platform, and I'd especially like to thank Jim Brandt, and the TPF team as a whole, for their patience, confidence and support throughout the course of this project and helping to get it across the finish line to a production, v1 release.

<br>

# Support

If you encounter any problems or have questions, feel free to reach out to me directly via E-Mail at vanstyn -at- intellitree.com. You can also always find me in #rapidapp on irc.perl.org.

Since this is quite a signification change, I'm sure that there will be issues to crop up. That's why I'm making myself directly available so I can address whatever does come up as quickly as possible.

Thank you to all, and I hope you enjoy the new system!

