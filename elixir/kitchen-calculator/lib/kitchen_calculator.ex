defmodule KitchenCalculator do
  @us_cup 240
  @us_fluid_ounce 30
  @us_teaspoon 5
  @us_tablespoon 15

  def get_volume({_, volume}) do
    volume
  end

  def to_milliliter({:cup, volume}) do
    {:milliliter, volume * @us_cup}
  end

  def to_milliliter({:fluid_ounce, volume}) do
    {:milliliter, volume * @us_fluid_ounce}
  end

  def to_milliliter({:teaspoon, volume}) do
    {:milliliter, volume * @us_teaspoon}
  end

  def to_milliliter({:tablespoon, volume}) do
    {:milliliter, volume * @us_tablespoon}
  end

  def to_milliliter({:milliliter, volume}) do
    {:milliliter, volume}
  end

  def from_milliliter({:milliliter, volume}, :milliliter) do
    {:milliliter, volume}
  end

  def from_milliliter({:milliliter, volume}, :cup) do
    {:cup, volume / @us_cup}
  end

  def from_milliliter({:milliliter, volume}, :fluid_ounce) do
    {:fluid_ounce, volume / @us_fluid_ounce}
  end

  def from_milliliter({:milliliter, volume}, :teaspoon) do
    {:teaspoon, volume / @us_teaspoon}
  end

  def from_milliliter({:milliliter, volume}, :tablespoon) do
    {:tablespoon, volume / @us_tablespoon}
  end

  def convert(volume_pair, unit) do
    volume_pair |> to_milliliter() |> from_milliliter(unit)
  end
end
