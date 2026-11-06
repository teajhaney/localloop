---
name: LocalLoop
colors:
  surface: '#fbf8fc'
  surface-dim: '#dcd9dd'
  surface-bright: '#fbf8fc'
  surface-container-lowest: '#ffffff'
  surface-container-low: '#f6f2f7'
  surface-container: '#f0edf1'
  surface-container-high: '#eae7eb'
  surface-container-highest: '#e4e1e6'
  on-surface: '#1b1b1e'
  on-surface-variant: '#464555'
  inverse-surface: '#303033'
  inverse-on-surface: '#f3f0f4'
  outline: '#777587'
  outline-variant: '#c7c4d8'
  surface-tint: '#4d44e3'
  primary: '#3525cd'
  on-primary: '#ffffff'
  primary-container: '#4f46e5'
  on-primary-container: '#dad7ff'
  inverse-primary: '#c3c0ff'
  secondary: '#006c49'
  on-secondary: '#ffffff'
  secondary-container: '#6cf8bb'
  on-secondary-container: '#00714d'
  tertiary: '#684000'
  on-tertiary: '#ffffff'
  tertiary-container: '#885500'
  on-tertiary-container: '#ffd4a4'
  error: '#ba1a1a'
  on-error: '#ffffff'
  error-container: '#ffdad6'
  on-error-container: '#93000a'
  primary-fixed: '#e2dfff'
  primary-fixed-dim: '#c3c0ff'
  on-primary-fixed: '#0f0069'
  on-primary-fixed-variant: '#3323cc'
  secondary-fixed: '#6ffbbe'
  secondary-fixed-dim: '#4edea3'
  on-secondary-fixed: '#002113'
  on-secondary-fixed-variant: '#005236'
  tertiary-fixed: '#ffddb8'
  tertiary-fixed-dim: '#ffb95f'
  on-tertiary-fixed: '#2a1700'
  on-tertiary-fixed-variant: '#653e00'
  background: '#fbf8fc'
  on-background: '#1b1b1e'
  surface-variant: '#e4e1e6'
typography:
  display-hero:
    fontFamily: Geist
    fontSize: 48px
    fontWeight: '600'
    lineHeight: 52px
    letterSpacing: -0.03em
  display-hero-mobile:
    fontFamily: Geist
    fontSize: 32px
    fontWeight: '600'
    lineHeight: 38px
    letterSpacing: -0.025em
  headline-lg:
    fontFamily: Geist
    fontSize: 30px
    fontWeight: '600'
    lineHeight: 36px
    letterSpacing: -0.025em
  headline-lg-mobile:
    fontFamily: Geist
    fontSize: 24px
    fontWeight: '600'
    lineHeight: 30px
    letterSpacing: -0.02em
  headline-md:
    fontFamily: Geist
    fontSize: 20px
    fontWeight: '600'
    lineHeight: 26px
    letterSpacing: -0.02em
  headline-sm:
    fontFamily: Geist
    fontSize: 16px
    fontWeight: '600'
    lineHeight: 22px
    letterSpacing: -0.015em
  body-lg:
    fontFamily: Geist
    fontSize: 17px
    fontWeight: '400'
    lineHeight: 26px
    letterSpacing: -0.01em
  body-md:
    fontFamily: Geist
    fontSize: 15px
    fontWeight: '400'
    lineHeight: 22px
    letterSpacing: -0.005em
  body-sm:
    fontFamily: Geist
    fontSize: 13px
    fontWeight: '400'
    lineHeight: 18px
    letterSpacing: 0em
  label-md:
    fontFamily: Geist
    fontSize: 13px
    fontWeight: '500'
    lineHeight: 16px
    letterSpacing: -0.005em
  label-sm:
    fontFamily: Geist
    fontSize: 11px
    fontWeight: '600'
    lineHeight: 14px
    letterSpacing: 0.02em
rounded:
  sm: 0.25rem
  DEFAULT: 0.5rem
  md: 0.75rem
  lg: 1rem
  xl: 1.5rem
  full: 9999px
spacing:
  gutter: 1.5rem
  gutter-sm: 1rem
  margin: 2rem
  margin-sm: 1rem
  margin-lg: 4rem
  space-xs: 0.25rem
  space-sm: 0.5rem
  space-md: 1rem
  space-lg: 1.5rem
  space-xl: 2rem
---

## Brand & Style

The design system embodies a modern, warm-minimalist aesthetic tailored for effortless local discovery and streamlined attendance. Inspired by the editorial restraint of Luma, the spatial warmth of Airbnb, and the functional precision of Linear, the interface operates as an invisible, high-clarity canvas for local experiences. 

### Key Characteristics
- **Personality**: Calm, intentional, community-rooted, and highly functional.
- **Target Audience**: Urban professionals, creatives, and community organizers who value quick discovery, frictionless free RSVPs, and clutter-free interfaces.
- **Visual Style**: Minimalist structure combined with soft physical tactility. High structural clarity is achieved through warm stone backdrops, crisp elevated card surfaces, hair-thin boundaries (1px subtle strokes), and intentional pops of vivid indigo for primary actions.
- **Design Principles**:
  - *Content Forward*: Imagery and typography communicate context before decorative elements do.
  - *Extreme Simplicity*: Free RSVP only; strictly remove transactional cruft, payment flows, reviews, and infinite feed noise.
  - *Tactile Precision*: Micro-interactions, soft ambient elevations, and subtle card hover-lifts give digital cards a real-world print feel.

## Colors

The color palette pairs warm stone neutrals with crisp white containers and an electric indigo accent, supported by unambiguous functional semantic tokens.

### Base Palettes
- **Canvas Base**: `#FAFAF9` (Stone 50) establishes a warm, approachable background that reduces eye strain compared to clinical white.
- **Surface Elevation**: `#FFFFFF` creates crisp figure-ground separation for cards, modals, sheets, and popovers.
- **Borders & Dividers**: `#E7E5E4` (Stone 200) provides 1px structural outlines with minimal visual weight.
- **Text & Hierarchy**:
  - `Text Primary`: `#18181B` (Zinc 900) for sharp, high-contrast readability.
  - `Text Secondary`: `#71717A` (Zinc 500) for timestamps, metadata, organizers, and secondary icons.
  - `Text Muted`: `#A1A1AA` (Zinc 400) for placeholders and subtle counters.

### Interactive & Semantic Roles
- **Primary Interactive (`#4F46E5`)**: Electric Indigo drives core actions: "RSVP Now", "Create Event", active tabs, and focused keyboard rings.
- **Semantic Success (`#10B981` / `#059669`)**: Emerald indicates registered states: "Going", "Confirmed", and open capacity statuses.
- **Semantic Warning (`#F59E0B` / `#D97706`)**: Amber highlights scarcity: "Few spots left", "Limited capacity".
- **Semantic Destructive (`#EF4444` / `#DC2626`)**: Rose/Red informs finality: "At Capacity", "Full", or "Event Canceled".

### Dark Mode Adaptability
All tokens map symmetrically: canvas adapts to `#09090B`, surfaces to `#18181B`, borders to `#27272A`, primary text shifts to `#FAFAF9`, and primary indigo softens slightly to `#6366F1` to balance contrast against dark backgrounds.

## Typography

The design system relies on `Geist` across all typographic roles to cultivate a clean, engineering-grade clarity balanced with contemporary editorial composure.

### Hierarchy & Tuning
- **Headings**: Set in Semibold (`600`) with negative letter spacing (`tracking-tight` from `-0.015em` to `-0.03em`). This maintains typographic density, preventing headlines from looking disconnected.
- **Body**: Regular (`400`) weight ensures prolonged legibility for descriptions, details, and guidelines. Generous line heights (`140%` to `150%`) ensure effortless horizontal scanning.
- **Labels & Badges**: Medium (`500`) and Semibold (`600`) weights anchor micro-metadata such as dates, calendar pills, attendee counts, and capacity states. Compact sizes use slight positive tracking (`0.02em`) to remain sharp at small dimensions.

## Layout & Spacing

Layouts follow an 8-point base grid inside a fixed-width container logic to prevent event discovery grids from ballooning on wide viewports.

### Grid & Max Dimensions
- **Max Width**: Central content containers cap at `1200px` (or `640px` for singular event detail/RSVP flows).
- **Desktop (>= 1024px)**: 12-column responsive layout or a 3-column event card grid. `margin-lg: 4rem` with `gutter: 1.5rem`.
- **Tablet (768px - 1023px)**: 2-column event grid. `margin: 2rem`, `gutter: 1.25rem`.
- **Mobile (< 768px)**: 1-column layout. Cards stack vertically. Outer margin resets to `margin-sm: 1rem` to maximize horizontal scan area.

### Internal Spacing Rhythm
- **Component Padding**: Standard card internal padding uses `space-lg` (24px) for desktop and `space-md` (16px) for compact viewports.
- **Micro Separation**: Detail items (e.g., date icon + text label) are separated by `space-sm` (8px). Consecutive metadata rows are spaced by `space-sm` to `space-md`.
- **Section Rhythm**: Sections on landing and discovery pages are separated by `space-xl` (32px) to 64px, preserving an uncluttered reading cadence.

## Elevation & Depth

Depth is tactile and subtle, relying on low-opacity ambient drop shadows combined with fine 1px strokes rather than heavy drop shadows.

### Elevation Tiers
1. **Flat / Base Level (`z0`)**:
   - The canvas surface `#FAFAF9`. No shadow.
2. **Raised Surfaces & Cards (`z1`)**:
   - Background: Pure White (`#FFFFFF`).
   - Border: `1px solid #E7E5E4`.
   - Shadow: `0 1px 3px 0 rgba(24, 24, 27, 0.04), 0 1px 2px -1px rgba(24, 24, 27, 0.03)`.
3. **Interactive Hover State (`z1-hover`)**:
   - Transform: `translateY(-2px)`.
   - Shadow: `0 8px 20px -4px rgba(24, 24, 27, 0.08), 0 4px 6px -2px rgba(24, 24, 27, 0.03)`.
   - Transition: `all 200ms cubic-bezier(0.16, 1, 0.3, 1)`.
4. **Overlays & Floating Modals (`z2`)**:
   - RSVP verification sheets, dialogs, and popovers.
   - Shadow: `0 20px 25px -5px rgba(24, 24, 27, 0.1), 0 8px 10px -6px rgba(24, 24, 27, 0.04)`.
   - Backdrop: `rgba(24, 24, 27, 0.4)` with `backdrop-filter: blur(4px)`.

## Shapes

The design system uses a balanced rounded geometry to create a friendly, modern consumer feel while maintaining typographic alignment.

### Geometry Rules
- **Event Containers & Cards**: Set to `rounded-lg` (16px / `1rem`). This cushions photography and softens container edges.
- **Inputs & Standard Action Buttons**: Set to `rounded` (10px to 12px) to match the internal radius of nested card components.
- **Badges, Status Chips, & Date Tags**: Fully rounded pills (`border-radius: 9999px`) provide contrast against the rectangular grid of cards.
- **Image Aspect Ratios**: Primary cover photos enforce a consistent `16:9` or `3:2` aspect ratio with matching top border-radii.

## Components

### 1. Navigation Bar
- **Structure**: Sticky minimal header (`height: 64px`), background `#FFFFFF` with `backdrop-filter: blur(8px)`, opacity `95%`, bottom border `1px solid #E7E5E4`.
- **Contents**: Left-aligned wordmark, center contextual filters (e.g., "This Weekend", "Tech", "Arts"), and right-aligned "Create Event" ghost button + "Sign In" or profile avatar pill.

### 2. Event Cards
- **Structure**: Top image container (`16:9` ratio, object-cover) followed by body padding (`space-md` to `space-lg`).
- **Header Elements**: Date badge overlay (white rounded pill with bold month/day in top-left).
- **Content Hierarchy**:
  - Capacity Status Badge (top-right or directly below image).
  - Event Title: `headline-md`, zinc-900, truncated at 2 lines.
  - Metadata row: Time icon + formatted range; Location pin + neighborhood name (e.g., "Mission District, SF").
  - Attendee Counter: `label-md`, zinc-500 (e.g., "42 going · 8 spots left").
- **Hover Behavior**: Card shifts `translateY(-2px)` with shadow elevation `z1-hover`.

### 3. Status Badges & Pills
- **Confirmed / Going**: Emerald background tint (`#ECFDF5`), emerald text (`#065F46`), border `1px solid #A7F3D0`.
- **Few Spots Left**: Amber background tint (`#FFFBEB`), amber text (`#92400E`), border `1px solid #FDE68A`.
- **Full / Closed**: Zinc background tint (`#F4F4F5`), zinc text (`#71717A`), border `1px solid #E4E4E7`.
- **Canceled**: Rose background tint (`#FEF2F2`), rose text (`#991B1B`), border `1px solid #FECACA`.

### 4. Action Buttons
- **Primary ("RSVP")**: Background `#4F46E5`, text `#FFFFFF`, font weight `500`, radius `10px`, hover `#4338CA`, active `#3730A3`. Focus ring: `2px solid #FFFFFF, 2px solid #4F46E5`.
- **Secondary / Subdued**: Background `#F4F4F5`, text `#18181B`, hover `#E4E4E7`.
- **Outline / Ghost**: Background transparent, border `1px solid #E7E5E4`, text `#18181B`, hover background `#F4F4F5`.
- **State: Attending**: When RSVP is confirmed, button transitions to emerald outline or fill: `#ECFDF5` background, `#059669` text, label: "✓ You're Going".

### 5. Form Fields & RSVP Sheet
- **Input Fields**: Single-row text/email inputs with `1px solid #E7E5E4`, radius `10px`, padding `10px 14px`, text `15px`. Active focus state transitions border to `#4F46E5` with `0 0 0 3px rgba(79, 70, 229, 0.12)`.
- **RSVP Flow**: A distraction-free bottom sheet (mobile) or centered card modal (desktop). Limited to name and email only. Disables checkout, ticketing options, or paid upsells.