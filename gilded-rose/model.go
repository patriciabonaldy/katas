package main

// Item interface.
type Item interface {
	UpdateItem()
	GetName() string
}

type item struct {
	name            string
	sellIn, quality int
}

func (i *item) UpdateItem() {
	i.updateQuality()
	i.updateSelling()
}

func (i *item) GetName() string {
	return i.name
}

func (i *item) updateQuality() {
	if i.quality > 0 {
		i.quality--
	}

	if i.sellIn < 0 && i.quality > 0 {
		i.quality--
	}
}

func (i *item) updateSelling() {
	i.sellIn--
}

type Brie struct {
	item
}

func (b *Brie) GetName() string {
	return b.name
}

func (b *Brie) UpdateItem() {
	b.updateQuality()
	b.updateSelling()
}

func (b *Brie) updateSelling() {
	b.sellIn--
}

func (b *Brie) updateQuality() {
	if b.quality < 50 {
		b.quality++
		if b.sellIn < 11 && b.quality < 50 {
			b.quality++
		}

		if b.sellIn < 6 && b.quality < 50 {
			b.quality++
		}
	}

	if b.sellIn < 0 && b.quality < 50 {
		b.quality++
	}
}

type Sulfuras struct {
	item
}

func (s *Sulfuras) GetName() string {
	return s.name
}

func (s *Sulfuras) UpdateItem() {
	s.updateQuality()
}

func (s *Sulfuras) updateQuality() {
	if s.quality < 50 {
		s.quality++
	}
}

type BackStage struct {
	item
}

func (b *BackStage) GetName() string {
	return b.name
}

func (b *BackStage) UpdateItem() {
	b.updateQuality()
	b.updateSelling()
}

func (b *BackStage) updateSelling() {
	b.sellIn--
}

func (b *BackStage) updateQuality() {
	if b.quality < 50 {
		b.quality++
		if b.sellIn < 11 && b.quality < 50 {
			b.quality++
		}

		if b.sellIn < 6 && b.quality < 50 {
			b.quality++
		}
	}

	if b.sellIn < 0 {
		b.quality++
	}
}