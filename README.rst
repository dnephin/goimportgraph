Go Import Graph
===============

Generate a ``.dot`` file from a go package import graph.

Install
-------

.. code::

    go get github.com/dnephin/goimportgraph


Example Usage
-------------

.. code::

    goimportgraph github.com/dnephin/goimportgraph > graph.dot
    dot -Tsvg -ograph.svg graph.dot
