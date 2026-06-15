---
Title: Goja D-Bus Module Intern Design Guide
Ticket: GOJA-DBUS-DESIGN
Status: active
Topics:
    - goja
    - dbus
    - design
DocType: index
Intent: long-term
Owners: []
RelatedFiles: []
ExternalSources:
    - ./sources/01-dbus.md
Summary: "Ticket workspace for an intern-facing Goja D-Bus native module design and implementation guide."
LastUpdated: 2026-06-15T17:45:00-04:00
WhatFor: "Use this ticket to review the proposed require(\"dbus\") module architecture and continue implementation planning."
WhenToUse: "When onboarding an intern, reviewing the D-Bus module design, or starting implementation work."
---

# Goja D-Bus Module Intern Design Guide

## Overview

This ticket contains a detailed design and implementation guide for a future `require("dbus")` native module for go-go-goja. The guide explains the JavaScript API, Go package boundaries, D-Bus concepts, runtime-owner safety rules, Promise and signal scheduling patterns, service export flow, policy model, testing strategy, and phased implementation plan.

## Key Links

- [Primary design guide](./design-doc/01-goja-d-bus-module-intern-design-and-implementation-guide.md)
- [Investigation diary](./reference/01-investigation-diary.md)
- [Imported source proposal](./sources/01-dbus.md)
- [Tasks](./tasks.md)
- [Changelog](./changelog.md)

## Status

Current status: **active**. The documentation deliverable is complete; implementation has not started.

## Topics

- goja
- dbus
- design

## Structure

- `design-doc/` — architecture, API, and implementation guide.
- `reference/` — investigation diary.
- `sources/` — imported source material, including `/tmp/dbus.md`.
- `tasks.md` — completed documentation tasks and implementation follow-ups.
- `changelog.md` — ticket change history.
