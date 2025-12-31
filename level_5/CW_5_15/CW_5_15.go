/*
Your company has just hired your friend from college, and you received a referral bonus. To celebrate, you decide to spend the bonus on beer cans and build the largest possible three-dimensional beer can pyramid.

Each level of the pyramid requires the square of the level number in beer cans:
	- Level 1 → 1 can;
	- Level 2 → 4 cans;
	- Level 3 → 9 cans;
	- Level 4 → 16 cans;
	- Level 5 → 25 cans;
	- and so on…

Write a function beeramid that returns the maximum number of complete levels you can build.

Parameters:
	- bonus - the amount of money you received as a referral bonus;
	- price - the price of a single beer can.

You can only build complete levels. Any remaining money that cannot be used to complete the next level should be ignored.

В вашей компании только что приняли на работу вашего друга из колледжа, и вы получили реферальный бонус. Чтобы отпраздновать это событие, вы решаете потратить бонус на покупку пивных банок и построить как можно большую трёхмерную пирамиду из банок.

Для каждого уровня пирамиды требуется количество банок, равное квадрату номера уровня:
	- 1-й уровень → 1 банка;
	- 2-й уровень → 4 банки;
	- 3-й уровень → 9 банок;
	- 4-й уровень → 16 банок;
	- 5-й уровень → 25 банок;
	- и так далее…

Необходимо написать функцию beeramid, которая возвращает максимальное количество полностью построенных уровней пирамиды.

Параметры:
	- bonus - сумма реферального бонуса;
	- price - цена одной пивной банки.

Строить можно только полные уровни. Оставшиеся деньги, которых недостаточно для следующего уровня, не учитываются.
*/

package kata

func Beeramid(bonus int, price float64) int {
	if bonus <= 0 || price <= 0 {
		return 0
	}

	totalCans := int(float64(bonus) / price)
	levels := 0
	requiredCans := 0

	for {
		nextLevel := levels + 1
		requiredCans += nextLevel * nextLevel

		if requiredCans > totalCans {
			break
		}

		levels = nextLevel
	}

	return levels
}