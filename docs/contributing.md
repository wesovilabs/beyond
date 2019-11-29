---
layout: page
title: Contributing
permalink: /contributors/
nav_order: 7
---

{: .text-green-300}
# Want to be part of Beyond?
{: .fs-9 }

{: .text-green-200}
Open source would die if there're were no contributors...
{: .fs-6 .fw-300 }

---

{: .text-green-300}
## Let's make other people know about Beyond

{: .text-yellow-300}
### Sharing Beyond via Social networks

It's really appreciated retrieving feedback from other people. And much more if they like what you
do and they share it with other people so 

<div class="socialme">
    <ul>
        <li class="twitter">
            <a href="https://twitter.com/intent/tweet?via={{site.data.social.twitter.username}}&url={{ site.data.social.twitter.url | uri_escape}}&text={{ site.data.social.twitter.message | uri_escape}}" target="_blank">
                {% include social/twitter.svg %}
            </a>
        </li>
        <li class="reddit">
            <a href="http://www.reddit.com/submit?url={{ site.data.social.reddit.message | uri_escape}}&title={{ site.data.social.reddit.title | uri_escape }}" target="_blank">
                {% include social/reddit.svg %}
            </a>
        </li>
        <li class="linkedin">
            <a href="https://www.linkedin.com/shareArticle?mini=true&url={{ site.data.social.linkedin.url | uri_escape}}&title={{ site.data.social.linkedin.title}}" target="_blank">
                {% include social/linkedin.svg %}
            </a>
        </li>
    </ul>
</div>




{: .text-yellow-300}
### Post articles about Beyond

Write technical guides or tutorials, and let others know how to use Beyond. I will be so happy to link your posts from
this site.
 
<div class="socialme">
 <ul>
     <li class="medium">
         <a href="{{ site.data.social.medium.url }}" target="_blank">
             {% include social/medium.svg %}
         </a>
     </li>
 </ul>
</div>

{: .text-yellow-300}
### Star Beyond on Github

It's awesome to find out that people recognize your work. If you like this project I ask you to star it.


<!-- Place this tag where you want the button to render. -->
<!-- Place this tag in your head or just before your close body tag. -->
<script async defer src="https://buttons.github.io/buttons.js"></script>
<a class="github-button" href="https://github.com/wesovilabs/beyond" data-color-scheme="no-preference: light; light: light; dark: dark;" data-icon="octicon-star" data-size="large" data-show-count="true" aria-label="Star wesovilabs/beyond on GitHub">Star</a>

{: .text-green-300}
## Do you want to contribute with your code?

Great projects are built when great developers work together. That's
the reason I encourage you to take part of Beyond!

Have a look at open issues and contribute with your code:

- [Fixing bugs](https://github.com/wesovilabs/beyond/projects/1)
- [Features on roadmap](https://github.com/wesovilabs/beyond/projects/2)
- [Implementing new features](https://github.com/wesovilabs/beyond/projects/4)]


{: .text-yellow-300}
### Working with Beyond code

#### Checkout the code

Fork the [Beyond repository](https://github.com/wesovilabs/beyond) and clone it locally 

```bash
git clone https://github.com/<user>/beyond.git
cd beyond
```

#### Setup Git hooks

```bash
make init
```

- commit-msg: It checks commit messages style.
- pre-commit: It formats your code.
- pre-push: It guarantees that your code pass tests and linter checks before being pushed.


#### Running the tests

Run tests to verify Beyond works

```bash
make test
```

Eventually you could check the test coverage with

```bash
make test-coverage
``` 

#### Check the source code and find potential optimizations

Beyond makes use of [golangci-lint](https://github.com/golangci/golangci-lint).

```bash
make lint
```

#### Build an executable

To build an executable for your current os, just run:

```bash
make run
```

Alternatively you could  easily generate Beyond executables for linux, darwin and 
windows at time by running.

```bash
make build-all
``` 

---

Don't hesitate to drop me an email at `ivan.corrales.solera@gmail.com` if you want to collaborate!

---

<div class="socialme">
    <ul>
        <li class="twitter">
            <a href="https://twitter.com/intent/tweet?via={{site.data.social.twitter.username}}&url={{ site.data.social.twitter.url | uri_escape}}&text={{ site.data.social.twitter.message | uri_escape}}" target="_blank">
                {% include social/twitter.svg %}
            </a>
        </li>
    </ul>
</div>
