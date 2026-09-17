

# Guerra Civis – Documento de Visão

## Introdução

### 1.1 Objetivo do sistema:
O objetivo principal do sistema Guerras Civis é proporcionar uma experiência de jogo multijogador online de estratégia e dedução social, inspirada nos clássicos "Cidade Dorme", "Detetive" e "Lobisomen". 

O sistema deve permitir que grupos de jogadores participem de partidas estruturadas em turnos (Noite e Dia), nas quais os participantes são secretamente divididos entre dois times rivais — Máfia e Civis — e recebem papéis com habilidades específicas. O objetivo de cada time é eliminar completamente o adversário por meio de ações noturnas e votações diurnas.

### 1.2 Escopo do desenvolvimento do sistema

  - Criação e gerenciamento de partidas
  - Distribuição secreta de papéis entre os jogadores.
  - Controle de fases do jogo (Noite e Dia)
  - Execução e resolução de habilidades noturnas
  - Sistema de votação diurna para eliminação de jogadores.
  - Verificação das condições de vitória.
  - Sistema de comunicação entre jogadores (chat geral, chat privado, chat de mortos).

## Visão geral do sistema

### 2.1 Visão e Objetivos do Sistema

A visão do sistema Guerras Civis é tornar-se a principal plataforma independente para a realização de partidas do clássico jogo de dedução social originário do Habbo Hotel, oferecendo uma experiência estável, justa, configurável e acessível.
O sistema busca preservar a essência do jogo original (estratégia, bluff, trabalho em equipe e tensão social).

Objetivos principais:

- Oferecer partidas fluidas e organizadas com controle claro de fases (Noite e Dia).
- Garantir distribuição justa e secreta de papéis.
- Criar uma base sólida e expansível para futuras melhorias (ranking, novos modos, etc.).

Objetivos secundários:

- Facilitar a organização de partidas por Hosts e comunidades.
- Servir como referência moderna do gênero Mafia/Werewolf no contexto brasileiro.

### 2.2 Contexto e Limite do Sistema

O sistema Guerras Civis opera como uma aplicação web e um aplicativo mobile multijogador em tempo real, focada exclusivamente na realização de partidas do jogo de dedução social.

Contexto:
- Os jogadores e Hosts interagem em salas virtuais de partida.
- O sistema gerencia todo o ciclo de vida de uma partida: criação, distribuição de papéis, execução de fases, resolução de ações e declaração de vencedor.

### 2.3 Estrutura Geral do sistema

O sistema é organizado em módulos principais que trabalham de forma integrada:

**Módulo de Gerenciamento de Partidas:**
- Responsável pela criação, entrada de jogadores, início e finalização das partidas.

**Módulo de Papéis e Habilidades:**
- Gerencia a distribuição secreta de papéis e a execução/resolução das habilidades noturnas.
  
**Módulo de Controle de Fases:**
- Controla a alternância entre Noite e Dia, temporizadores e anúncios oficiais.
  
**Módulo de Votação e Eliminação:**
- Processa as votações diurnas e aplica as eliminações.
  
**Módulo de Comunicação:**
- Gerencia os diferentes canais de chat.
  
**Módulo de Estado e Persistência**
- Mantém o estado atual da partida, status dos jogadores e logs básicos.
