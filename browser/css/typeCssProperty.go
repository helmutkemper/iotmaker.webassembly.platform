package css

// https://developer.mozilla.org/en-US/docs/Web/CSS/Reference#index

type Property string

func (e Property) String() string {
	return string(e)
}

const (

	// KPropertyAccentColor
	//
	//  https://developer.mozilla.org/en-US/docs/Web/CSS/accent-color
	//
	// English:
	//
	// Sets the accent color for user-interface controls generated by some elements.
	//
	//  Values:
	//
	//    auto: Represents a UA-chosen color, which should match the accent color of the platform, if any.
	//
	//    <color>: Specifies the color to be used as the accent color.
	//
	//  Notes:
	//    * Browsers that support accent-color currently apply it to the following HTML elements:
	//      <input type="checkbox">
	//      <input type="radio">
	//      <input type="range">
	//      <progress>
	//
	// Português:
	//
	// Define a cor de destaque para os controles da interface do usuário gerados por alguns elementos.
	//
	//  Valores:
	//
	//    auto: Representa uma cor escolhida pelo UA, que deve corresponder à cor de destaque da plataforma, se houver.
	//
	//    <color>: Especifica a cor a ser usada como cor de destaque.
	//
	//  Notas:
	//    * Os navegadores que suportam a cor de destaque atualmente a aplicam aos seguintes elementos HTML:
	//      <input type="checkbox">
	//      <input type="radio">
	//      <input type="range">
	//      <progress>
	KPropertyAccentColor Property = "accent-color"

	// KPropertyAlignContent
	//
	//  https://developer.mozilla.org/en-US/docs/Web/CSS/align-content
	//
	// English:
	//
	// Sets the distribution of space between and around content items along a flexbox's cross-axis or a grid's block
	// axis.
	//
	//  Values:
	//
	//    start: The items are packed flush to each other against the start edge of the alignment container in the
	//      cross axis.
	//
	//    end: The items are packed flush to each other against the end edge of the alignment container in the cross axis.
	//
	//    flex-start: The items are packed flush to each other against the edge of the alignment container depending on
	//      the flex container's cross-start side. This only applies to flex layout items. For items that are not
	//      children of a flex container, this value is treated like start.
	//
	//    flex-end: The items are packed flush to each other against the edge of the alignment container depending on the
	//      flex container's cross-end side. This only applies to flex layout items. For items that are not children of a
	//      flex container, this value is treated like end.
	//
	//    center: The items are packed flush to each other in the center of the alignment container along the cross axis.
	//
	//    normal: The items are packed in their default position as if no align-content value was set.
	//
	//    baseline, first baseline, last baseline: Specifies participation in first- or last-baseline alignment: aligns
	//      the alignment baseline of the box's first or last baseline set with the corresponding baseline in the shared
	//      first or last baseline set of all the boxes in its baseline-sharing group.
	//
	//    space-between: The items are evenly distributed within the alignment container along the cross axis.
	//      The spacing between each pair of adjacent items is the same. The first item is flush with the start edge of
	//      the alignment container in the cross axis, and the last item is flush with the end edge of the alignment
	//      container in the cross axis.
	//
	//    space-around: The items are evenly distributed within the alignment container along the cross axis.
	//      The spacing between each pair of adjacent items is the same. The empty space before the first and after the
	//      last item equals half of the space between each pair of adjacent items.
	//
	//    space-evenly: The items are evenly distributed within the alignment container along the cross axis.
	//      The spacing between each pair of adjacent items, the start edge and the first item, and the end edge and the
	//      last item, are all exactly the same.
	//
	//    stretch: If the combined size of the items along the cross axis is less than the size of the alignment
	//      container, any auto-sized items have their size increased equally (not proportionally), while still respecting
	//      the constraints imposed by max-height/max-width (or equivalent functionality), so that the combined size
	//      exactly fills the alignment container along the cross axis.
	//
	//    safe: Used alongside an alignment keyword. If the chosen keyword means that the item overflows the alignment
	//      container causing data loss, the item is instead aligned as if the alignment mode were start.
	//
	//    unsafe: Used alongside an alignment keyword. Regardless of the relative sizes of the item and alignment
	//      container and whether overflow which causes data loss might happen, the given alignment value is honored.
	//
	//  Notes:
	//    * This property has no effect on single line flex containers (i.e. ones with flex-wrap: nowrap).
	//
	// Português:
	//
	// Define a distribuição do espaço entre e ao redor dos itens de conteúdo ao longo do eixo cruzado de um flexbox ou
	// do eixo de bloco de uma grade.
	//
	//  Valores:
	//
	//    start: Os itens são empacotados alinhados uns com os outros contra a borda inicial do contêiner de alinhamento
	//      no eixo cruzado.
	//
	//    end: Os itens são embalados alinhados uns com os outros contra a borda final do contêiner de alinhamento no
	//      eixo transversal.
	//
	//    flex-start: Os itens são embalados alinhados uns com os outros contra a borda do contêiner de alinhamento,
	//      dependendo do lado de partida cruzada do contêiner flexível. Isso se aplica apenas a itens de layout flexível.
	//      Para itens que não são filhos de um contêiner flexível, esse valor é tratado como start.
	//
	//    flex-end: Os itens são embalados alinhados uns com os outros contra a borda do contêiner de alinhamento,
	//      dependendo do lado da extremidade transversal do contêiner flexível. Isso se aplica apenas a itens de layout
	//      flexível. Para itens que não são filhos de um contêiner flexível, esse valor é tratado como end.
	//
	//    center: Os itens são embalados alinhados uns com os outros no centro do contêiner de alinhamento ao longo do
	//      eixo transversal.
	//
	//    normal: Os itens são empacotados em sua posição padrão como se nenhum valor de alinhamento de conteúdo
	//      fosse definido.
	//
	//    baseline, first baseline, last baseline: Especifica a participação no alinhamento da primeira ou última linha
	//      de base: alinha a linha de base de alinhamento do primeiro ou último conjunto de linha de base da caixa com a
	//      linha de base correspondente no primeiro ou último conjunto de linha de base compartilhado de todas as caixas
	//      em seu grupo de compartilhamento de linha de base.
	//
	//    space-between: Os itens são distribuídos uniformemente dentro do contêiner de alinhamento ao longo do eixo
	//      cruzado. O espaçamento entre cada par de itens adjacentes é o mesmo. O primeiro item está alinhado com a borda
	//      inicial do contêiner de alinhamento no eixo cruzado e o último item está alinhado com a borda final do
	//      contêiner de alinhamento no eixo cruzado.
	//
	//    space-around: Os itens são distribuídos uniformemente dentro do contêiner de alinhamento ao longo do eixo
	//      cruzado. O espaçamento entre cada par de itens adjacentes é o mesmo. O espaço vazio antes do primeiro e
	//      depois do último item é igual a metade do espaço entre cada par de itens adjacentes.
	//
	//    space-evenly: Os itens são distribuídos uniformemente dentro do contêiner de alinhamento ao longo do eixo
	//      cruzado. O espaçamento entre cada par de itens adjacentes, a borda inicial e o primeiro item, e a borda final
	//      e o último item, são exatamente iguais.
	//
	//    stretch: Se o tamanho combinado dos itens ao longo do eixo cruzado for menor que o tamanho do contêiner de
	//      alinhamento, quaisquer itens de tamanho automático terão seu tamanho aumentado igualmente (não
	//      proporcionalmente), respeitando as restrições impostas por max-heightmax-width (ou funcionalidade
	//      equivalente), para que o tamanho combinado preencha exatamente o contêiner de alinhamento ao longo do
	//      eixo cruzado.
	//
	//    safe: Usado junto com uma palavra-chave de alinhamento. Se a palavra-chave escolhida significar que o item
	//      estoura no contêiner de alinhamento causando perda de dados, o item é alinhado como se o modo de alinhamento
	//      fosse iniciado.
	//
	//    unsafe: Usado junto com uma palavra-chave de alinhamento. Independentemente dos tamanhos relativos do item e do
	//      contêiner de alinhamento e se o estouro que causa perda de dados pode ocorrer, o valor de alinhamento
	//      fornecido é respeitado.
	//
	//  Notas:
	//    * Esta propriedade não tem efeito em contêineres flexíveis de linha única (ou seja, aqueles com flex-wrap:
	//      nowrap).
	KPropertyAlignContent Property = "align-content"

	// KPropertyAlignItems
	//
	//  https://developer.mozilla.org/en-US/docs/Web/CSS/align-items
	//
	// English:
	//
	// Sets the align-self value on all direct children as a group. In Flexbox, it controls the alignment of items on the Cross Axis. In Grid Layout, it controls the alignment of items on the Block Axis within their grid area.
	//
	//  Values:
	//
	//    normal: The effect of this keyword is dependent of the layout mode we are in:
	//      * In absolutely-positioned layouts, the keyword behaves like start on replaced absolutely-positioned boxes,
	//        and as stretch on all other absolutely-positioned boxes.
	//      * In static position of absolutely-positioned layouts, the keyword behaves as stretch.
	//      * For flex items, the keyword behaves as stretch.
	//      * For grid items, this keyword leads to a behavior similar to the one of stretch, except for boxes with an
	//        aspect ratio or an intrinsic sizes where it behaves like start.
	//      * The property doesn't apply to block-level boxes, and to table cells.
	//
	//    flex-start: The cross-start margin edges of the flex items are flushed with the cross-start edge of the line.
	//
	//    flex-end: The cross-end margin edges of the flex items are flushed with the cross-end edge of the line.
	//
	//    center: The flex items' margin boxes are centered within the line on the cross-axis. If the cross-size of an
	//      item is larger than the flex container, it will overflow equally in both directions.
	//
	//    start: The items are packed flush to each other toward the start edge of the alignment container in the
	//      appropriate axis.
	//
	//    end: The items are packed flush to each other toward the end edge of the alignment container in the
	//      appropriate axis.
	//
	//    self-start: The items are packed flush to the edge of the alignment container of the start side of the item,
	//      in the appropriate axis.
	//
	//    self-end: The items are packed flush to the edge of the alignment container of the end side of the item, in
	//      the appropriate axis.
	//
	//    baseline, first baseline, last baseline: All flex items are aligned such that their flex container baselines
	//      align. The item with the largest distance between its cross-start margin edge and its baseline is flushed
	//      with the cross-start edge of the line.
	//
	//    stretch: Flex items are stretched such that the cross-size of the item's margin box is the same as the line
	//      while respecting width and height constraints.
	//
	//    safe: Used alongside an alignment keyword. If the chosen keyword means that the item overflows the alignment
	//      container causing data loss, the item is instead aligned as if the alignment mode were start.
	//
	//    unsafe: Used alongside an alignment keyword. Regardless of the relative sizes of the item and alignment
	//      container and whether overflow which causes data loss might happen, the given alignment value is honored.
	//
	// Português:
	//
	// Define o valor align-self em todos os filhos diretos como um grupo. No Flexbox, ele controla o alinhamento dos
	// itens no Cross Axis. No Layout de Grade, ele controla o alinhamento dos itens no Eixo do Bloco dentro de sua área
	// de grade.
	//
	//  Valores:
	//
	//    normal: O efeito desta palavra-chave depende do modo de layout em que estamos:
	//      * Em layouts absolutamente posicionados, a palavra-chave se comporta como start em caixas absolutamente
	//        posicionadas substituídas e como stretch em todas as outras caixas absolutamente posicionadas.
	//      * Na posição estática de layouts absolutamente posicionados, a palavra-chave se comporta como esticar.
	//      * Para itens flexíveis, a palavra-chave se comporta como esticar.
	//      * Para itens de grade, essa palavra-chave leva a um comportamento semelhante ao de esticar, exceto para
	//        caixas com uma proporção ou tamanhos intrínsecos onde se comporta como início.
	//      * A propriedade não se aplica a caixas de nível de bloco e a células de tabela.
	//
	//    flex-start: As bordas da margem de início cruzado dos itens flexíveis são niveladas com a borda de início
	//      cruzado da linha.
	//
	//    flex-end: As bordas da margem cruzada dos itens flexíveis são niveladas com a borda cruzada da linha.
	//
	//    center: As caixas de margem dos itens flexíveis são centralizadas dentro da linha no eixo cruzado.
	//      Se o tamanho cruzado de um item for maior que o contêiner flexível, ele transbordará igualmente em ambas as
	//      direções.
	//
	//    start: Os itens são empacotados alinhados uns com os outros em direção à borda inicial do contêiner de
	//      alinhamento no eixo apropriado.
	//
	//    end: Os itens são embalados alinhados uns com os outros em direção à borda final do contêiner de alinhamento
	//      no eixo apropriado.
	//
	//    self-start: Os itens são empacotados rente à borda do contêiner de alinhamento do lado inicial do item, no
	//      eixo apropriado.
	//
	//    self-end: Os itens são embalados rente à borda do contêiner de alinhamento do lado final do item, no eixo
	//      apropriado.
	//
	//    baseline, first baseline, last baseline: Todos os itens flexíveis são alinhados de forma que suas linhas de
	//      base do contêiner flexível se alinhem. O item com a maior distância entre sua borda de margem de início
	//      cruzado e sua linha de base é nivelado com a borda de início cruzado da linha.
	//
	//    stretch: Os itens flexíveis são esticados de forma que o tamanho cruzado da caixa de margem do item seja o
	//      mesmo da linha, respeitando as restrições de largura e altura.
	//
	//    safe: Usado junto com uma palavra-chave de alinhamento. Se a palavra-chave escolhida significar que o item
	//      estoura no contêiner de alinhamento causando perda de dados, o item é alinhado como se o modo de alinhamento
	//      fosse iniciado.
	//
	//    unsafe: Usado junto com uma palavra-chave de alinhamento. Independentemente dos tamanhos relativos do item e do
	//      contêiner de alinhamento e se o estouro que causa perda de dados pode ocorrer, o valor de alinhamento
	//      fornecido é respeitado.
	KPropertyAlignItems Property = "align-items"

	// KPropertyAlignSelf
	//
	//  https://developer.mozilla.org/en-US/docs/Web/CSS/align-self
	//
	// English:
	//
	// Overrides a grid or flex item's align-items value. In Grid, it aligns the item inside the grid area.
	// In Flexbox, it aligns the item on the cross axis.
	//
	//  Values:
	//
	//    auto: Computes to the parent's align-items value.
	//
	//    normal: The effect of this keyword is dependent of the layout mode we are in:
	//      * In absolutely-positioned layouts, the keyword behaves like start on replaced absolutely-positioned boxes,
	//        and as stretch on all other absolutely-positioned boxes.
	//      * In static position of absolutely-positioned layouts, the keyword behaves as stretch.
	//      * For flex items, the keyword behaves as stretch.
	//      * For grid items, this keyword leads to a behavior similar to the one of stretch, except for boxes with an
	//        aspect ratio or an intrinsic sizes where it behaves like start.
	//      * The property doesn't apply to block-level boxes, and to table cells.
	//
	//    self-start: Aligns the items to be flush with the edge of the alignment container corresponding to the item's
	//      start side in the cross axis.
	//
	//    self-end: Aligns the items to be flush with the edge of the alignment container corresponding to the item's
	//      end side in the cross axis.
	//
	//    flex-start: The cross-start margin edge of the flex item is flushed with the cross-start edge of the line.
	//
	//    flex-end: The cross-end margin edge of the flex item is flushed with the cross-end edge of the line.
	//
	//    center: The flex item's margin box is centered within the line on the cross-axis. If the cross-size of the item
	//      is larger than the flex container, it will overflow equally in both directions.
	//
	//    baseline, first baseline, last baseline: Specifies participation in first- or last-baseline alignment: aligns
	//      the alignment baseline of the box's first or last baseline set with the corresponding baseline in the shared
	//      first or last baseline set of all the boxes in its baseline-sharing group. The fallback alignment for first
	//      baseline is start, the one for last baseline is end.
	//
	//    stretch: If the combined size of the items along the cross axis is less than the size of the alignment
	//      container and the item is auto-sized, its size is increased equally (not proportionally), while still
	//      respecting the constraints imposed by max-height/max-width (or equivalent functionality), so that the
	//      combined size of all auto-sized items exactly fills the alignment container along the cross axis.
	//
	//    safe: If the size of the item overflows the alignment container, the item is instead aligned as if the
	//      alignment mode were start.
	//
	//    unsafe: Regardless of the relative sizes of the item and alignment container, the given alignment value is
	//      honored.
	//
	//  Notes:
	//    * The property doesn't apply to block-level boxes, or to table cells.
	//      If a flexbox item's cross-axis margin is auto, then align-self is ignored.
	//
	// Português:
	//
	// Substitui o valor align-items de uma grade ou item flexível. Em Grade, alinha o item dentro da área da grade.
	// No Flexbox, alinha o item no eixo cruzado.
	//
	//  Valores:
	//
	//    auto: Calcula para o valor align-items do pai.
	//
	//    normal: O efeito desta palavra-chave depende do modo de layout em que estamos:
	//      * Em layouts absolutamente posicionados, a palavra-chave se comporta como start em caixas absolutamente
	//        posicionadas substituídas e como stretch em todas as outras caixas absolutamente posicionadas.
	//      * Na posição estática de layouts absolutamente posicionados, a palavra-chave se comporta como esticar.
	//      * Para itens flexíveis, a palavra-chave se comporta como esticar.
	//      * Para itens de grade, essa palavra-chave leva a um comportamento semelhante ao de esticar, exceto para
	//        caixas com uma proporção ou tamanhos intrínsecos onde se comporta como início.
	//      * A propriedade não se aplica a caixas de nível de bloco e a células de tabela.
	//
	//    self-start: Alinha os itens a serem alinhados com a borda do contêiner de alinhamento correspondente ao lado
	//      inicial do item no eixo cruzado.
	//
	//    self-end: Alinha os itens a serem alinhados com a borda do contêiner de alinhamento correspondente ao lado
	//      final do item no eixo cruzado.
	//
	//    flex-start: A borda da margem de início cruzado do item flexível é nivelada com a borda de início cruzado
	//      da linha.
	//
	//    flex-end: A borda da margem cruzada do item flexível é nivelada com a borda cruzada da linha.
	//
	//    center: A caixa de margem de itens flexíveis é centralizada dentro da linha no eixo cruzado.
	//      Se o tamanho cruzado do item for maior que o contêiner flexível, ele transbordará igualmente em ambas as
	//      direções.
	//
	//    baseline, first baseline, last baseline: Especifica a participação no alinhamento da primeira ou última linha
	//      de base: alinha a linha de base de alinhamento do primeiro ou último conjunto de linha de base da caixa com a
	//      linha de base correspondente no primeiro ou último conjunto de linha de base compartilhado de todas as caixas
	//      em seu grupo de compartilhamento de linha de base. O alinhamento de fallback para a primeira linha de base é
	//      start, o da última linha de base é end.
	//
	//    stretch: Se o tamanho combinado dos itens ao longo do eixo transversal for menor que o tamanho do contêiner de
	//      alinhamento e o item for dimensionado automaticamente, seu tamanho será aumentado igualmente (não
	//      proporcionalmente), respeitando as restrições impostas por max-heightmax-width (ou funcionalidade
	//      equivalente), para que o tamanho combinado de todos os itens de dimensionamento automático preencha exatamente
	//      o contêiner de alinhamento ao longo do eixo cruzado.
	//
	//    safe: Se o tamanho do item ultrapassar o contêiner de alinhamento, o item será alinhado como se o modo de
	//      alinhamento tivesse sido iniciado.
	//
	//    unsafe: Independentemente dos tamanhos relativos do item e do contêiner de alinhamento, o valor de alinhamento
	//      fornecido é respeitado.
	//
	//  Notas:
	//    * A propriedade não se aplica a caixas de nível de bloco ou a células de tabela.
	//      Se a margem do eixo cruzado de um item do flexbox for automática, align-self será ignorado.
	KPropertyAlignSelf Property = "align-self"

	// KPropertyAlignTracks
	//
	//  https://developer.mozilla.org/en-US/docs/Web/CSS/align-tracks
	//
	// English:
	//
	// Sets the alignment in the masonry axis for grid containers that have masonry in their block axis.
	//
	// The property can take a single value, in which case all tracks are aligned in the same way. If a list of values is used then the first value applies to the first track in the grid axis, the second to the next and so on.
	//
	// If there are fewer values than tracks, the last value is used for all remaining tracks. If there are more values than tracks, any additional values are ignored.
	//
	//  Values:
	//
	//    start: The items are packed flush to each other toward the start edge of the alignment container in the
	//      masonry axis.
	//
	//    end: The items are packed flush to each other toward the end edge of the alignment container in the
	//      masonry axis.
	//
	//    center: The items are packed flush to each other toward the center of the alignment container along the
	//      masonry axis.
	//
	//    normal: Acts as start.
	//
	//    baseline | first baseline | last baseline: Specifies participation in first- or last-baseline alignment: aligns
	//      the alignment baseline of the box's first or last baseline set with the corresponding baseline in the shared
	//      first or last baseline set of all the boxes in its baseline-sharing group. The fallback alignment for first
	//      baseline is start, the one for last baseline is end.
	//
	//    space-between: The items are evenly distributed within the alignment container along the masonry axis.
	//      The spacing between each pair of adjacent items is the same. The first item is flush with the main-start edge,
	//      and the last item is flush with the main-end edge.
	//
	//    space-around: The items are evenly distributed within the alignment container along the masonry axis.
	//      The spacing between each pair of adjacent items is the same. The empty space before the first and after the
	//      last item equals half of the space between each pair of adjacent items.
	//
	//    space-evenly: The items are evenly distributed within the alignment container along the masonry axis.
	//      The spacing between each pair of adjacent items, the main-start edge and the first item, and the main-end edge
	//      and the last item, are all exactly the same.
	//
	//    stretch: The items stretch along the masonry axis to fill the content box. Items with definite size do not
	//      stretch.
	//
	// Português:
	//
	// Define o alinhamento no eixo de alvenaria para contêineres de grade que possuem alvenaria em seu eixo de bloco.
	//
	// A propriedade pode ter um único valor, nesse caso todas as trilhas são alinhadas da mesma maneira. Se uma lista de
	// valores for usada, o primeiro valor será aplicado à primeira faixa no eixo da grade, o segundo à próxima e assim
	// por diante.
	//
	// Se houver menos valores do que faixas, o último valor será usado para todas as faixas restantes. Se houver mais
	// valores do que faixas, quaisquer valores adicionais serão ignorados.
	//
	//  Valores:
	//
	//    start: Os itens são embalados alinhados uns com os outros em direção à borda inicial do recipiente de
	//      alinhamento no eixo de alvenaria.
	//
	//    end: Os itens são embalados alinhados uns com os outros em direção à borda final do recipiente de alinhamento
	//      no eixo de alvenaria.
	//
	//    center: Os itens são embalados alinhados uns com os outros em direção ao centro do recipiente de alinhamento ao
	//      longo do eixo de alvenaria.
	//
	//    normal: Atua como início.
	//
	//    baseline | first baseline | last baseline: Especifica a participação no alinhamento da primeira ou última linha
	//      de base: alinha a linha de base de alinhamento do primeiro ou último conjunto de linha de base da caixa com a
	//      linha de base correspondente no primeiro ou último conjunto de linha de base compartilhado de todas as caixas
	//      em seu grupo de compartilhamento de linha de base. O alinhamento de fallback para a primeira linha de base é
	//      start, o da última linha de base é end.
	//
	//    space-between: Os itens são distribuídos uniformemente dentro do contêiner de alinhamento ao longo do eixo de
	//      alvenaria. O espaçamento entre cada par de itens adjacentes é o mesmo. O primeiro item está alinhado com a
	//      borda de início principal e o último item está alinhado com a borda de extremidade principal.
	//
	//    space-around: Os itens são distribuídos uniformemente dentro do contêiner de alinhamento ao longo do eixo de
	//      alvenaria. O espaçamento entre cada par de itens adjacentes é o mesmo. O espaço vazio antes do primeiro e
	//      depois do último item é igual a metade do espaço entre cada par de itens adjacentes.
	//
	//    space-evenly: Os itens são distribuídos uniformemente dentro do contêiner de alinhamento ao longo do eixo de
	//      alvenaria. O espaçamento entre cada par de itens adjacentes, a borda inicial principal e o primeiro item, e a
	//      borda final principal e o último item, são exatamente iguais.
	//
	//    stretch: Os itens se estendem ao longo do eixo de alvenaria para preencher a caixa de conteúdo. Itens com
	//      tamanho definido não estica.
	KPropertyAlignTracks Property = "align-tracks"

	// KPropertyAll
	//
	//  https://developer.mozilla.org/en-US/docs/Web/CSS/all
	//
	// English:
	//
	// Resets all of an element's properties except unicode-bidi, direction, and CSS Custom Properties. It can set
	// properties to their initial or inherited values, or to the values specified in another cascade layer or stylesheet
	// origin.
	//
	// The all property is specified as one of the CSS global keyword values. Note that none of these values affect the
	// unicode-bidi and direction properties.
	//
	//  Values:
	//
	//    initial: Specifies that all the element's properties should be changed to their initial values.
	//
	//    inherit: Specifies that all the element's properties should be changed to their inherited values.
	//
	//    unset: Specifies that all the element's properties should be changed to their inherited values if they inherit
	//      by default, or to their initial values if not.
	//
	//    revert: Specifies behavior that depends on the stylesheet origin to which the declaration belongs:
	//      * If the rule belongs to the author origin, the revert value rolls back the cascade to the user level, so
	//        that the specified values are calculated as if no author-level rules were specified for the element.
	//        For purposes of revert, the author origin includes the Override and Animation origins.
	//      * If the rule belongs to the user origin, the revert value rolls back the cascade to the user-agent level,
	//        so that the specified values are calculated as if no author-level or user-level rules were specified for
	//        the element.
	//      * If the rule belongs to the user-agent origin, the revert value acts like unset.
	//
	//    revert-layer: Specifies that all the element's properties should roll back the cascade to a previous cascade
	//      layer, if one exists. If no other cascade layer exists, the element's properties will roll back to the
	//      matching rule, if one exists, in the current layer or to a previous style origin.
	//
	// Português:
	//
	// Redefine todas as propriedades de um elemento, exceto unicode-bidi, direction e CSS Custom Properties. Ele pode
	// definir propriedades para seus valores iniciais ou herdados, ou para os valores especificados em outra camada em
	// cascata ou origem de folha de estilo.
	//
	// A propriedade all é especificada como um dos valores de palavras-chave globais do CSS. Observe que nenhum desses
	// valores afeta as propriedades unicode-bidi e direction.
	//
	//  Valores:
	//
	//    initial: Especifica que todas as propriedades do elemento devem ser alteradas para seus valores iniciais.
	//
	//    inherit: Especifica que todas as propriedades do elemento devem ser alteradas para seus valores herdados.
	//
	//    unset: Especifica que todas as propriedades do elemento devem ser alteradas para seus valores herdados, se
	//      herdarem por padrão, ou para seus valores iniciais, se não forem.
	//
	//    revert: Especifica o comportamento que depende da origem da folha de estilo à qual a declaração pertence:
	//      * Se a regra pertencer à origem do autor, o valor de reversão reverte a cascata para o nível do usuário, para
	//        que os valores especificados sejam calculados como se nenhuma regra no nível do autor tivesse sido
	//        especificada para o elemento. Para fins de reversão, a origem do autor inclui as origens Override e
	//        Animation.
	//      * Se a regra pertencer à origem do usuário, o valor de reversão reverte a cascata para o nível de agente do
	//        usuário, de modo que os valores especificados sejam calculados como se nenhuma regra de nível de autor ou
	//        nível de usuário tivesse sido especificada para o elemento.
	//      * Se a regra pertence à origem do agente do usuário, o valor de reversão age como não definido.
	//
	//    revert-layer: Especifica que todas as propriedades do elemento devem reverter a cascata para uma camada de
	//      cascata anterior, se houver. Se nenhuma outra camada em cascata existir, as propriedades do elemento serão
	//      revertidas para a regra de correspondência, se existir, na camada atual ou em uma origem de estilo anterior.
	KPropertyAll Property = "all"

	// KPropertyAngle
	//
	//  https://developer.mozilla.org/en-US/docs/Web/CSS/angle
	//
	// English:
	//
	// Data type represents an angle value expressed in degrees, gradians, radians, or turns. It is used, for example,
	// in <gradient>s and in some transform functions.
	//
	// The <angle> data type consists of a <number> followed by one of the units listed below. As with all dimensions,
	// there is no space between the unit literal and the number. The angle unit is optional after the number 0.
	//
	// Optionally, it may be preceded by a single + or - sign. Positive numbers represent clockwise angles, while negative
	// numbers represent counterclockwise angles. For static properties of a given unit, any angle can be represented by
	// various equivalent values. For example, 90deg equals -270deg, and 1turn equals 4turn. For dynamic properties, like
	// when applying an animation or transition, the effect will nevertheless be different.
	//
	//  Units:
	//
	//    deg: Represents an angle in degrees. One full circle is 360deg. Examples: 0deg, 90deg, 14.23deg.
	//
	//    grad: Represents an angle in gradians. One full circle is 400grad. Examples: 0grad, 100grad, 38.8grad.
	//
	//    rad: Represents an angle in radians. One full circle is 2π radians which approximates to 6.2832rad. 1rad is
	//      180/π degrees. Examples: 0rad, 1.0708rad, 6.2832rad.
	//
	//    turn: Represents an angle in a number of turns. One full circle is 1turn. Examples: 0turn, 0.25turn, 1.2turn.
	//
	// Português:
	//
	// O tipo de dados representa um valor de ângulo expresso em graus, gradianos, radianos ou voltas. É usado, por
	// exemplo, em <gradient>s e em algumas funções de transformação.
	//
	// O tipo de dados <angle> consiste em um <number> seguido por uma das unidades listadas abaixo. Como em todas as
	// dimensões, não há espaço entre o literal da unidade e o número. A unidade angular é opcional após o número 0.
	//
	// Opcionalmente, pode ser precedido por um único sinal + ou -. Números positivos representam ângulos no sentido
	// horário, enquanto números negativos representam ângulos no sentido anti-horário. Para propriedades estáticas de
	// uma determinada unidade, qualquer ângulo pode ser representado por vários valores equivalentes. Por exemplo,
	// 90 graus é igual a -270 graus e 1 volta é igual a 4 voltas. Para propriedades dinâmicas, como ao aplicar uma
	// animação ou transição, o efeito será diferente.
	//
	//  Unidades:
	//
	//    deg: Representa um ângulo em graus. Um círculo completo é 360deg. Exemplos: 0 graus, 90 graus, 14,23 graus.
	//
	//    grad: Representa um ângulo em gradianos. Um círculo completo é 400 grad. Exemplos: 0 grad, 100 grad, 38,8 grad.
	//
	//    rad: Representa um ângulo em radianos. Um círculo completo é 2π radianos que se aproxima de 6,2832rad. 1rad é
	//      180π graus. Exemplos: 0rad, 1,0708rad, 6,2832rad.
	//
	//    turn: Representa um ângulo em um número de voltas. Um círculo completo é 1 volta. Exemplos: 0 volta,
	//      0,25 volta, 1,2 volta.
	KPropertyAngle Property = "angle"
)

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
