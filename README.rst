Wikipedia game solver
-----------------------

Find the shortest path between any two wikipedia pages.

Idea
====

In a single pass, import page titles and links from wikipedia's xml dump.
Given two pages, find the shortest path between them using A* search.

Should not require enough memory to load all of the data. To do this, store
the mapping of page titles and links in boltdb.

Usage
======

Dependencies
============

http://github.com/boltdb/bolt - used to persist data on disk
