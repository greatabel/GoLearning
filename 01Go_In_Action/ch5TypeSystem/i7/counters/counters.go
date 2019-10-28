package counters

type alertCounter int


func New(value int) alertCounter {
    return alertCounter(value)
}