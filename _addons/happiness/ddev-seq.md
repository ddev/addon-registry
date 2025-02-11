---
title: happiness/ddev-seq
github_url: https://github.com/happiness/ddev-seq
description: "This is a DDEV add-on for Seq, a self-hosted search, analysis, and alerting server built for structured logs and traces."
user: happiness
repo: ddev-seq
categories:
  - community
created_at: 2025-02-10
updated_at: 2025-02-10
stars: 0
---

# ddev-seq

This is a DDEV add-on for [Seq](https://datalust.co/seq), a self-hosted search,
analysis, and alerting server built for structured logs and traces.

## Add-on installation

```
ddev add-on get happiness/ddev-seq
```

## Drupal integration

Install the modules [OpenTelemetry](https://www.drupal.org/project/opentelemetry) and
[OpenTelemetry Log](https://www.drupal.org/project/otlog). The *OpenTelemetry* module
comes with it's own log implementation but I have not managed to get that working
properly so for now I recommend using the *OpenTelemetry Log* module instead.

```
composer require drupal/opentelemetry drupal/otlog
```

Configure the *OpenTelemetry* module at `/admin/config/development/opentelemetry`
and add the **Endpoint** `http://seq/ingest/otlp`.

That's it. Open `http://seq.<DDEV_NAME>.ddev.site:8501/` to start exploring
your logs and traces.

### Use OpenTelementry bundled log module

If you want to use the log module bundled with the [OpenTelemetry](https://www.drupal.org/project/opentelemetry)
module you need to apply the patch from [the following issue](https://www.drupal.org/project/opentelemetry/issues/3505594)
to get the log message body into Seq.

## Tracing

The following example of how to do your own tracing in your custom code is
taken from the [OpenTelemetry](https://www.drupal.org/project/opentelemetry)
module page with some minor modifications.

```php
$tracer_provider = \Drupal::service('opentelemetry.tracer_provider');
$tracer = $tracer_provider->getTracer('drupal_test');
$main_span = $tracer->spanBuilder('My custom operation')->startSpan();

// Make an external API call.
$api_span = $tracer->spanBuilder("Coindesk API call")->startSpan();
$data = json_decode(
    file_get_contents('https://api.coindesk.com/v1/bpi/currentprice.json')
);
$api_span->end();

// Set attributes and put an event to the main span.
$bots = 3;
$main_span->setAttribute('Bitcoin value', $data->bpi->USD->rate_float);
$main_span->addEvent('Starting bots tuning', ['bots available' => $bots]);

for ($i = 1; $i <= $bots; $i++) {
    $inner_span = $tracer->spanBuilder("Tuning bot $i")->startSpan();
    // Do some internal business logic.
    usleep(rand(200000, 500000));
    $raised[$i] = rand(0, 100);
    $inner_span->addEvent("Bot $i raised money!", ['amount' => $raised[$i]]);
    usleep(rand(100000, 200000));
    $inner_span->end();
}

// Do some more stuff and finalize the main span.
usleep(rand(200000, 500000));
$main_span->addEvent('We got richer!', ['raised' => array_sum($raised)]);
$main_span->setAttribute('raised_details', $raised);
$main_span->end();
```
