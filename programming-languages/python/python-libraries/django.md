# [Django](https://www.djangoproject.com/)

## Notes

- Make migration (where `./manage` is a script that wraps over `python manage.py`):
  1. Change model. Add/remove fields etc.
  2. Run `./manage makemigrations` to create migration from changed models. (can run `./manage makemigrations -n <migration-name>` to give it custom name).
  3. Migration file is created. Run `./manage migrate` to apply
- Run tests:
  - `./manage test` = run all tests
  - `./manage test -k path.to.test` = run specific test. can also run some function from the test inside the file with another `.`.
- Migrate down/up: `./manage migrate <app> <migration-number>`
- Make empty migration: `./manage makemigrations --empty <name-of-migration> <app>`

## Links

- [Profiling Django with DTrace and cProfile (2019)](https://wiedi.frubar.net/blog/2019/11/18/django-performance/)
- [Django’s CBVs were a mistake (2012)](https://lukeplant.me.uk/blog/posts/djangos-cbvs-were-a-mistake/)
- [Testing Django Migrations (2019)](https://sobolevn.me/2019/10/testing-django-migrations)
- [HN: Django 3](https://news.ycombinator.com/item?id=21681732)
- [Viewflow](https://github.com/viewflow/viewflow) - Reusable workflow library for Django.
- [cookiecutter-django-rest](https://github.com/agconti/cookiecutter-django-rest) - Build best practiced apis fast with Python3.
- [rules](https://github.com/dfunckt/django-rules) - Awesome Django authorization, without the database.
- [Django REST framework](https://github.com/encode/django-rest-framework) - Awesome web-browsable Web APIs.
- [Awesome Django](https://github.com/wsvincent/awesome-django)
- [Introducing Django (2005)](https://simonwillison.net/2005/Jul/17/django/)
- [Django 3.1 (2020)](https://www.djangoproject.com/weblog/2020/aug/04/django-31-released/) ([HN](https://news.ycombinator.com/item?id=24048046))
- [Surviving Django (if you care about databases) (2020)](https://www.varrazzo.com/blog/2020/07/25/surviving-django/) ([HN](https://news.ycombinator.com/item?id=24074520)) ([Lobsters](https://lobste.rs/s/l7dqrf/surviving_django_if_you_care_about))
- [Django Views — The Right Way](https://spookylukey.github.io/django-views-the-right-way/) - Opinionated guide on how to write views in Django. ([Lobsters](https://lobste.rs/s/4n63nn/django_views_right_way))
- [Django Middleware](https://www.reddit.com/r/django/comments/hms3n6/django_middleware/)
- [Django Async: What's new and what's next? (2020)](https://deepsource.io/blog/django-async-support/) ([HN](https://news.ycombinator.com/item?id=24160608))
- [Django Redis](https://github.com/jazzband/django-redis) - Redis cache backend for Django.
- [Django styleguide](https://github.com/HackSoftware/Django-Styleguide)
- [Deploy Machine Learning Models with Django](https://www.deploymachinelearning.com/) ([Code](https://github.com/pplonski/my_ml_service))
- [Async Views in Django 3.1 (2020)](https://testdriven.io/blog/django-async-views/) ([HN](https://news.ycombinator.com/item?id=24423637))
- [Build a simple Hacker News clone using Django 3 (2020)](https://www.youtube.com/watch?v=292GB6snFYo)
- [Building a Django app with data access control in 30 minutes (2020)](https://www.osohq.com/post/django-access-control) ([Lobsters](https://lobste.rs/s/tpoc8j/building_django_app_with_data_access))
- [Django Waffle](https://github.com/django-waffle/django-waffle) - Feature flipper for Django.
- [Cookiecutter Django](https://github.com/pydanny/cookiecutter-django) - Framework for jump starting production-ready Django projects quickly.
- [SaaS Pegasus](https://www.saaspegasus.com/) - Django-Powered SaaS Template. ([Docs](https://docs.saaspegasus.com/))
- [Django Filter](https://github.com/carltongibson/django-filter) - Reusable Django application allowing users to declaratively add dynamic QuerySet filtering from URL parameters. ([Docs](https://django-filter.readthedocs.io/en/master/))
- [Django Best Practices: Security (2020)](https://learndjango.com/tutorials/django-best-practices-security)
- [LearnDjango](https://learndjango.com/) - Django Tutorials.
- [Django Dynamic Fixture](https://github.com/paulocheque/django-dynamic-fixture) - Complete library to create dynamic model instances for testing purposes.
- [Two Scoops of Django 3.x: Best Practices for the Django Web Framework](https://www.feldroy.com/products/two-scoops-of-django-3-x)
- [Speed Up Your Django Tests book](https://gumroad.com/l/suydt)
- [django-read-only](https://github.com/adamchainz/django-read-only) - Disable Django database writes.
- [Simple JWT](https://github.com/SimpleJWT/django-rest-framework-simplejwt) - JSON Web Token authentication plugin for the Django REST Framework.
- [drf-yasg](https://github.com/axnsan12/drf-yasg) - Generate real Swagger/OpenAPI 2.0 specifications from a Django Rest Framework API.
- [Graphene-Django](https://github.com/graphql-python/graphene-django) - Integrate GraphQL into your Django project.
- [Optimizing Postgres full text search in Django (2019)](https://hodovi.ch/blog/optimizing-postgres-full-text-search-django/)
- [Understand Group by in Django with SQL (2020)](https://hakibenita.com/django-group-by-sql)
- [Running Django + React service by Cloud Run (2020)](http://djkooks.github.io/gcp-django-deploy-cloudrun)
- [Django Doctor](https://django.doctor/) - Django GitHub PR bot that suggest the fix.
- [Django’s Structure – A Heretic’s Eye View](https://djangobook.com/mdj2-django-structure/)
- [Django Book](https://djangobook.com/) - Python Django Tutorials.
- [Blazing fast tests in Django (2018)](https://dizballanze.com/django-blazing-fast-tests/)
- [Django project optimization guide (2017)](https://dizballanze.com/django-project-optimization-part-1/)
- [Django Ninja](https://github.com/vitalik/django-ninja) - Fast Django REST Framework.
- [django-guardian](https://github.com/django-guardian/django-guardian) - Per object permissions for Django.
- [Docker-Compose for Django and React with Nginx Reverse-Proxy and Let's Encrypt (2020)](https://saasitive.com/tutorial/docker-compose-django-react-nginx-let-s-encrypt/) ([HN](https://news.ycombinator.com/item?id=25169155))
- [Django Smartmin](https://github.com/nyaruka/smartmin) - Admin-like utility for users.
- [Deploying Django to AWS ECS with Terraform](https://github.com/testdrivenio/django-ecs-terraform)
- [Django Slick Reporting](https://github.com/ra-systems/django-slick-reporting) - Powerful and Efficient reporting engine with Charting capabilities.
- [Dataclasses serializer](https://github.com/oxan/djangorestframework-dataclasses) - Dataclasses serializer for Django REST framework.
- [django-floppyforms](https://github.com/jazzband/django-floppyforms) - Full control of form rendering in the templates.
- [DIY Django and React Boilerplate for SaaS](https://github.com/saasitive/django-react-boilerplate) ([Web](https://saasitive.com/)) ([HN](https://news.ycombinator.com/item?id=25517226))
- [django-unicorn](https://github.com/adamghill/django-unicorn) - Provides a way to use backend Django code and regular Django templates to create interactive experiences without investing in a separate frontend framework. ([Web](https://www.django-unicorn.com/))
- [Hotwire + Django](https://github.com/hotwire-django/hotwire-django) - Meta package to combine turbo-django and stimulus-django.
- [Django migrations without downtimes (2015)](http://pankrat.github.io/2015/django-migrations-without-downtimes/)
- [wemake-django-template](https://github.com/wemake-services/wemake-django-template) - Bleeding edge django template focused on code quality and security.
- [Django Activity Stream](https://github.com/justquick/django-activity-stream) - Way of creating activities generated by the actions on your site.
- [Django antipatterns](https://www.django-antipatterns.com/) ([Code](https://github.com/hapytex/django-antipatterns))
- [Turbo for Django](https://github.com/hotwire-django/turbo-django) - Integration of Hotwire Turbo with Django.
- [Modern Django: A Guide on How to Deploy Django-based Web Applications](https://github.com/djstein/modern-django)
- [Django Algolia Integration](https://github.com/algolia/algoliasearch-django)
- [Docker, Django, Traefik, and IntercoolerJS: My go-to stack for building a SaaS (2021)](https://www.simplecto.com/docker-django-traefik-intercoolerjs-is-my-stack-for-2021/) ([HN](https://news.ycombinator.com/item?id=25973242))
- [Django 3.2 (2021)](https://www.djangoproject.com/weblog/2021/apr/06/django-32-released/) ([HN](https://news.ycombinator.com/item?id=26710013))
- [Constance](https://github.com/jazzband/django-constance) - Dynamic Django settings.
- [Rapid Prototyping with Django, htmx, and Tailwind CSS (2021)](https://testdriven.io/blog/django-htmx-tailwind/)
