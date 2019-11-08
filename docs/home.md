---
layout: default
title: Home
description: "Goa is a Golang Library that provides AOP"
nav_order: 1
permalink: /home
---

{: .text-green-300}
# Golang Oriented to Aspects (Goa)
{: .fs-9 }

{: .text-green-200}
A Golang library that drives us to the AOP paradigm.
{: .fs-6 .fw-300 }

[Get started now](#getting-started){: .btn .btn-primary .fs-5 .mb-4 .mb-md-0 .mr-2 } [View it on GitHub](https://github.com/wesovilabs/goa){: .btn .fs-5 .mb-4 .mb-md-0 }

---

{: .text-blue-300}
## Getting started

{: .text-blue-100}
### Installation

To make use of Goa, we just need to add the Goa dependency to our go.mod file. 


**go.mod**

```text
module github.com/wesovilabs/goa-examples/greetings
...
require github.com/wesovilabs/goa master
...
```

Available Goa releases can be found [here](https://github.com/wesovilabs/goa/releases)

---

## License
Goa is distributed by an [MIT license](https://github.com/wesovilabs/goa/tree/master/LICENSE.md).

## Contributing
When contributing to this repository, please first discuss the change you wish to make via issue, email, or any other 
method with the owners of this repository before making a change. 

Read more about becoming a contributor in our [GitHub repository](https://github.com/wesovilabs/goa/tree/master/contributing.md).

#### Thank you to the contributors of Goa!

<ul class="list-style-none">
{% for contributor in site.github.contributors %}
  <li class="d-inline-block mr-1">
     <a href="{{ contributor.html_url }}"><img src="{{ contributor.avatar_url }}" width="32" height="32" alt="{{ contributor.login }}"/></a>
  </li>
{% endfor %}
</ul>

## Code of Conduct
Goa is committed to fostering a welcoming community.

View our Code of Conduct on our [GitHub repository](https://github.com/wesovilabs/goa/tree/master/CODE_OF_CONDUCT.md).
