<template>
	<div class="container">
		<div class="block">
			<el-date-picker v-model="value1" type="daterange" range-separator="至" start-placeholder="开始日期"
				end-placeholder="结束日期">
			</el-date-picker>
		</div>
	</div>
</template>

<script>
	export default {
		name: "Date",
		components: {},
		data() {
			return {
				shortcuts: [{
					text: '最近一周',
					value: (() => {
						const end = new Date()
						const start = new Date()
						start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
						return [start, end]
					})(),
				}, {
					text: '最近一个月',
					value: (() => {
						const end = new Date()
						const start = new Date()
						start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
						return [start, end]
					})(),
				}, {
					text: '最近三个月',
					value: (() => {
						const end = new Date()
						const start = new Date()
						start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
						return [start, end]
					})(),
				}],
				value1: '',
				value2: ''
			};
		}
	};
</script>

<style>
</style>
