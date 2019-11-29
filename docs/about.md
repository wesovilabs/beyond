---
layout: page
title: About
permalink: /about/
nav_order: 8
---

{: .text-green-300}
# About Beyond
{: .fs-9 }

Beyond (Go Oriented to Aspects) is a Go library whose main beyondls are:
- It provides Go developers with a mechanism to **code under AOP paradigm**.
- You can build useful advices and **share with other developers**. at time you can also take advantages of existing 
ones.
- Beyond usage is really straightforward.
 
---

{: .text-green-300}
## About me

Hey folks! My name is Ivan and I could describe myself like an open source lover who enjoys doing some researches and 
sharing the knowledge with others.

Any comment, feedback or suggestion will be appreciated. Feel free to drop me an email at `ivan.corrales.solera@gmail.com`
or reach me at any of these social networks. 


{% include share.html %} 

---

{: .text-green-300}
## License
Beyond is distributed by an [MIT license](https://github.com/wesovilabs/beyond/tree/master/LICENSE.md).

{: .text-green-300}
## Contributing
When contributing to this repository, please first discuss the change you wish to make via issue, email, or any other 
method with the owners of this repository before making a change. 

Read more about becoming a contributor in our [Beyond repository](https://github.com/wesovilabs/beyond/tree/master/CONTRIBUTING.md).

{: .text-yellow-300}
### Thank you to the contributors of Beyond!

<ul class="list-style-none">
{% for contributor in site.github.contributors %}
  <li class="d-inline-block mr-1">
     <a href="{{ contributor.html_url }}"><img src="{{ contributor.avatar_url }}" width="32" height="32" alt="{{ contributor.login }}"/></a>
  </li>
{% endfor %}
</ul>

{: .text-green-300}
## Code of Conduct
Beyond is committed to fostering a welcoming community.

View Beyond Code of Conduct on  [Beyond repository](https://github.com/wesovilabs/beyond/tree/master/CODE_OF_CONDUCT.md).

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
