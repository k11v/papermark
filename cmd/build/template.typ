// model / bibliography

// model / cite

// model / document

// model / emph

// model / enum

#set enum(
    numbering: "1)",
    indent: 1.25cm,
)

// model / figure

// model / footnote

// model / heading
// TODO: https://forum.typst.app/t/how-to-conditionally-set-heading-block-parameters/2080/3

#set heading(numbering: "1.")
#show heading: set text(hyphenate: false)

#show heading.where(level: 1): set align(center)
#show heading.where(level: 1): set block(above: 15pt, below: 10pt)
#show heading.where(level: 1): set text(size: 16pt)
#show heading.where(level: 1): upper

#show heading.where(level: 2): set align(center)
#show heading.where(level: 2): set block(above: 15pt, below: 10pt)
#show heading.where(level: 2): set text(size: 14pt)

#show selector.or(..(3, 4, 5, 6).map(i => heading.where(level: i))): set align(left)
#show selector.or(..(3, 4, 5, 6).map(i => heading.where(level: i))): set block(above: 15pt, below: 10pt)
#show selector.or(..(3, 4, 5, 6).map(i => heading.where(level: i))): set text(size: 14pt)
#show selector.or(..(3, 4, 5, 6).map(i => heading.where(level: i))): emph

// model / link

// model / list

#set list(
    marker: [--],
    indent: 1.25cm,
)

// model / numbering

// model / outline

// model / par

#set par(
    leading: 1.05em, // Microsoft Word's 1.5
    spacing: 1.05em, // Microsoft Word's 1.5
    justify: true,
    first-line-indent: (amount: 1.25cm, all: true),
)

// model / parbreak

// model / quote

// model / ref

// model / strong

// model / table

// model / terms

// text / highlight

// text / linebreak

// text / lorem

// text / lower

// text / overline

// text / raw

#show raw: set text(font: "Courier New")

// text / smallcaps

// text / smartquote

// text / strike

// text / sub

// text / super

// text / text

#set text(
    font: "Times New Roman",
    size: 14pt,
    lang: "ru",
    hyphenate: auto,
)

// text / underline

// text / upper

// math / math.equation

// layout / align

// layout / alignment

// layout / angle

// layout / block

// layout / box

// layout / colbreak

// layout / columns

// layout / direction

// layout / fraction

// layout / grid

// layout / h

// layout / hide

// layout / layout

// layout / length

// layout / measure

// layout / move

// layout / pad

// layout / page

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

// layout / pagebreak

// layout / place

// layout / ratio

// layout / relative

// layout / repeat

// layout / rotate

// layout / scale

// layout / skew

// layout / stack

// layout / v

// visualize / circle

// visualize / color

// visualize / curve

// visualize / ellipse

// visualize / gradient

// visualize / image

// visualize / line

// visualize / path

// visualize / polygon

// visualize / rect

// visualize / square

// visualize / stroke

// visualize / tiling
