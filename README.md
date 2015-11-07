# DoGen

DoGen is allows you to generate documents on the fly through an API through a predefined template.

## Goal

The goal of DoGen is to simplify document generation. We want to make it easy for DoGen consumers to generate a document through a uniform interface, which can then be outputted in a number of formats.

The generation of these documents should all go through a predefined template, which will be passed to the renderer. The renderer parses these templates and outputs them in a desired output format.

**Note: for now, we will focus on PDF generation**

## Future

### Templates

At the moment, templates are tightly coupled with document output. This should be decoupled in the near future.

### Documents

For now, the only available document is the PDF type. This reflects in the document interface, which is heavily focused on the underlying PDF library. This should be abstracted and have a more general interface.

This reflects back on the templates which currently depend on a lot of the document internals. We'll start tearing this out and make a template less aware of what happens underneath the document methods.

### CLI

At the moment, there is only a web interface to generate documents. There should be a CLI instance as well.
