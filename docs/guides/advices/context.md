---
layout: post
parent: Advices
grand_parent: Guides & Tutorials
title: Goa context
nav_order: 4
---

# Goa context

- **Pkg()**: It returns the package of the intercepted function.

- **Function()**: It returns the name of the intercepted function.

- **Type()**: It returns the object of the intercepted method.

- **In() \*Args**: It returns a pointer of Args. It contains details from the function arguments.

- **Out() \*Args**: It returns a pointer of Args. It contains details from the function results.

- **Set(key string, value interface{})**: It permits to set custom data to the Goa context.

- **Get(key string) interface{}**: It returns a data that is stored in the Goa Context.
