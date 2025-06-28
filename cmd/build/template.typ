#set page(
    paper: "a4",
    margin: (
        top: 2cm,
        right: 1cm,
        bottom: 2cm,
        left: 3cm,
    ),
    numbering: "1",
)

#set text(
    font: "Times New Roman",
    size: 14pt,
    lang: "ru",
    hyphenate: auto,
)

#show raw: set text(font: "Courier New")

#set par(
    leading: 1.05em, // Microsoft Word's 1.5
    spacing: 1.05em, // Microsoft Word's 1.5
    justify: true,
    first-line-indent: (amount: 1.25cm, all: true),
)

#set list(
    marker: [--],
    indent: 1.25cm,
)

#set enum(
    numbering: "1.",
    indent: 1.25cm,
)
