

# Guerra Civis – Documento de Visão

## Introdução

### 1.1 Objetivo do sistema:
O objetivo principal do sistema Guerras Civis é proporcionar uma experiência de jogo multijogador online de estratégia e dedução social, inspirada nos clássicos "Cidade Dorme", "Detetive" e "Lobisomen", que substituirá a condunção manual por uma plataforma web/mobile automatizada. 

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

### 2.4 Características do Usuário

**Jogador Comum:**
- Familiarizado com jogos sociais/de dedução (tipo Detetive/Lobisomem).
- Não possui, necessariamente, conhecimento técnico avançado — a interface deve ser simples e guiada.
- Espera uma experiência com baixa curva de aprendizado, dado que já conhece as regras do jogo tradicional; a dificuldade deve estar na estratégia do jogo, não na usabilidade do sistema.
- Acesso predominante via dispositivos móveis.

## Requisitos

### 3.1 Requisitos por Subsistema / Componente

**1. Módulo de Gerenciamento de Partidas**
  - Criar, iniciar, reiniciar e finalizar partidas.
  - Controlar entrada e saída de jogadores.
  - Definir número mínimo e máximo de participantes.
  - Gerenciar status geral da partida (Aguardando, Em andamento, Finalizada).

**2. Módulo de Papéis e Habilidades**
  -  Distribuir papéis de forma aleatória, secreta e equilibrada.
  - Executar e resolver habilidades noturnas com ordem de prioridade.
  - Impedir uso de habilidades por jogadores mortos.
  - Enviar feedback privado ao jogador sobre o resultado de sua ação.

**3. Módulo Controle de Fases**
  - Alternar automaticamente entre as fases Noite e Dia.
  - Controlar temporizadores configuráveis de cada fase.
  - Emitir anúncios oficiais de início e fim de fase.
  - Bloquear ações incompatíveis com a fase atual.

**4. Módulo de Votação e Eliminação**
  - Permitir votação apenas de jogadores vivos durante o Dia.
  - Contabilizar votos e aplicar regras de empate.
  - Eliminar o jogador mais votado e atualizar seu status.
  - Revelar o papel do eliminado, conforme configuração.

**5. Módulo de Comunicação**
  - Gerenciar chat geral (restrito apenas ao dia).
  - Gerenciar chat de mortos/espectadores.
  - Gerenciar chat privado entre jogadores.
  - Impedir comunicação indevida entre vivos e mortos.

**7. Módulo de Estado e Persistência**
  - Manter o estado atual de todos os jogadores e da partida.
  - Registrar log básico de ações e eventos.
  - Suportar reconexão de jogadores mantendo seu estado.
  - Verificar automaticamente as condições de vitória.

### 3.2 Requisitos Funcionais 

**Gerenciamento de Partidas**
- RF01: Permitir criação de uma nova partida. (com ou sem senha)
- RF02: Definir um número mínimo e máximo de jogadores. (mínimo 12, máximo 18)
- RF03: Permitir entrada e saída de jogadores antes do início.
- RF04: Iniciar a partida apenas quando atingir o número mínimo de jogadores e o Host confirmar.
- RF05: Permitir reinício da partida ou criação de nova após o fim.

**Distribuição de Papéis**
- RF06: Distribuir papéis de forma aleatória e secreta entre os jogadores.
- RF07: Garantir equilíbrio entre times.
- RF08: Enviar o papel de forma privada para cada jogador.
- RF09: Permitir configuração de quais papéis estarão ativos na partida (lista de papéis disponíveis).

**Fases do Jogo**
- RF10: Alternar automaticamente entre as fases: Noite e Dia.
- RF11: Controlar o tempo de cada fase (configurável).
- RF12: Anunciar o início e o fim de cada fase para todos os jogadores.
- RF13: Bloquear ações indevidas durante cada fase

**Ações Noturnas**
- RF14: Permitir que jogadores usem suas habilidades apenas durante a noite
- RF15: Processar as ações na ordem de prioridade correta
- RF16: Resolver conflito de ações (ex: médico salvou o alvo do assassino)
- RF17: Impedir que um jogador morto use as habilidades.
- RF18: Enviar feedback privado ao jogador sobre o resultado da sua ação (quando aplicável)
- RF19: Bloquear chat geral durante a noite

**Fase do Dia**

- RF20: Permitir discussão livre entre os jogadores vivos.
- RF21: Iniciar processo de votação para eliminação.
- RF22: Permitir que cada jogador vivo vote em outro jogador vivo.
- RF23: Contabilizar os votos e eliminar o jogador mais votado.
- RF24: Anunciar o resultado da votação e a morte do jogador eliminado.

**Sistema de Vida e Eliminação**
- RF25: Manter status de cada jogador: Vivo / Morto / Espectator
- RF26: Permitir que jogadores mortos observem a partida (modo espectator), mas sem interagir com jogadores vivos.
- RF27: Revelar o papel do jogador morto.

**Condição de Vitória**
- RF28: Verificar automaticamente as condições de vitória ao final de cada fase.
- RF29: Mafia vence se o número de máfias for igual ou maior que o número de civis vivos.
- RF30: Civis vence se todos da máfia forem eliminados.
- RF31: Anunciar o time vencedor e revelar todos os papés ao final da partida.

**Comunicação e Inteface**
- RF32: Chat privado entre jogadores.
- RF33: Chat de mortos/espectatores separado.
- RF34: Interface clara mostrando: atual fase, tempo restante e lista de jogadores vivos
- RF35: Explicação das regras do jogo e das habilidades acessível ao jogador a qualquer momento.


**Requisitos de Qualidade**

| Categoria        | Requisito                                                                                                                               |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| Usabilidade      | A interface deve ser clara e intuitiva, permitindo que jogadores iniciantes compreendam a fase atual e suas ações em menos de 1 minuto. |
| Desempenho       | O sistema deve processar ações noturnas e votações em menos de 2 segundos, mesmo com 20 jogadores.                                      |
| Confiabilidade   | O sistema deve manter o estado da partida em caso de queda de conexão de até 30 segundos (reconexão).                                   |
| Disponibilidade  | O sistema deve estar disponível 99% do tempo durante horários de pico.                                                                  |
| Segurança        | Papéis e ações noturnas devem ser visíveis apenas para o jogador autorizado.                                                            |
| Manutenibilidade | O sistema deve ser modular, permitindo adição de novos papéis sem impacto nos módulos existentes.                                       |
| Escalabilidade   | O sistema deve suportar múltiplas partidas simultâneas (mínimo 50 partidas concorrentes na versão inicial).                             |

### 3.3 Interfaces
  
- **Interface do jogador:** Visualização da fase atual, tempo restante, lista de jogadores, chat e botões de ação (habilidades e votação)
- **Interface de Comunicação em Tempo Real:** Utilização de WebSocket para sincronização de estado, chat, e ações entre cliente e servidor.  
