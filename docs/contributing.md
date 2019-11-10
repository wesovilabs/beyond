---
layout: page
title: Contributing
permalink: /contributors/
nav_order: 3
---

{: .text-green-300}
# Want to be part of Goa?
{: .fs-9 }

{: .text-green-200}
Open source would die if there're were no contributors...
{: .fs-6 .fw-300 }

---

{: .text-green-300}
## Let's make other people know about Goa

{: .text-yellow-300}
### Sharing Goa via Social networks

When you spend time on doing open source projects, one of the most appreciated There's nothing most appreciated than retrieve feedback from other people. And much more if they like what you
do and they let you know and share with other people so 

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
            <a href="https://www.linkedin.com/shareArticle?mini=true&url={{ site.data.social.linkedin.url | uri_escape}}&title={{ site.data.social.linkedin.title | uri_escape}}" target="_blank">
                {% include social/linkedin.svg %}
            </a>
        </li>
    </ul>
</div>




{: .text-yellow-300}
### Post articles about Goa

Write technical guides or tutorials, and let others know how to use Goa. I will be so happy to link your posts from
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
### Star Goa on Github

It's awesome to find out that people recognize your work. If you
really like this project I will invite you to star it.


<!-- Place this tag where you want the button to render. -->
<!-- Place this tag in your head or just before your close body tag. -->
<script async defer src="https://buttons.github.io/buttons.js"></script>
<a class="github-button" href="https://github.com/wesovilabs/goa" data-color-scheme="no-preference: light; light: light; dark: dark;" data-icon="octicon-star" data-size="large" data-show-count="true" aria-label="Star wesovilabs/goa on GitHub">Star</a>

{: .text-green-300}
## Do you want to contribute with your code?

Great projects are build when great developers work together. That's
the reason I encourage you to take part of Goa!

Have a look at open issues and contribute with your code:

- [Fixing bugs](https://github.com/wesovilabs/goa/projects/1)
- [Implementing new features](https://github.com/wesovilabs/goa/projects/2)


{: .text-yellow-300}
### Working with Goa code


#### Checkout the code

Fork the [Goa repository](https://github.com/wesovilabs/goa) and clone it locally 

```bash
git clone https://github.com/<user>/goa.git
cd goa
```

#### Setup Git hooks

```bash
make init
```

- commit-msg: It checks commit messages style.
- pre-commit: It formats your code.
- pre-push: It guarantees that your code pass tests and linter checks before being pushed.


#### Running the tests

Run tests to verify Goa works as expected

```bash
make test
```

eventually you could check the test coverage with

```bash
make test-coverage
``` 

#### Check the source code and find potential optimizations

Goa makes use of [golangci-lint](https://github.com/golangci/golangci-lint).

```bash
make lint
```

#### Build an executable

To build an executable for your current os, just run:

```bash
make run
```

Alternatively you could  easily generate Goa executables for linux, darwin and 
windows at time by running.

```bash
make build-all
``` 

---

> Don't hesitate to drop me an email at `ivan.corrales.solera@gmail.com` if you have any doubt!
