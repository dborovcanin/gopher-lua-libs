local bit = require 'bit'
local assert = require 'assert'
local require = require 'require'

function Test_and(t)
    local result, err = bit.band(1, 0)
    require:NoError(t, err)
    assert:Equal(t, 0, result)
    result, err = bit.band(5, 6)
    require:NoError(t, err)
    assert:Equal(t, 4, result)
end

function Test_or(t)
    local result, err = bit.bor(1, 0)
    require:NoError(t, err)
    assert:Equal(t, 1, result)
    result, err = bit.bor(5, 6)
    require:NoError(t, err)
    assert:Equal(t, 7, result)
end

function Test_xor(t)
    local result, err = bit.bxor(1, 0)
    require:NoError(t, err)
    assert:Equal(t, 1, result)
    result, err = bit.bxor(5, 6)
    require:NoError(t, err)
    assert:Equal(t, 3, result)
end

function Test_left_shift(t)
    local result, err = bit.lshift(1, 0)
    require:NoError(t, err)
    assert:Equal(t, 1, result)
    result, err = bit.lshift(0xFF, 8)
    require:NoError(t, err)
    assert:Equal(t, 65280, result)
end

function Test_right_shift(t)
    local result, err = bit.rshift(42, 2)
    require:NoError(t, err)
    assert:Equal(t, 10, result)
    result, err = bit.rshift(0xFF, 4)
    require:NoError(t, err)
    assert:Equal(t, 15, result)
end

function Test_not(t)
    local result, err = bit.bnot(65536)
    require:NoError(t, err)
    assert:Equal(t, 4294901759, result)
    result, err = bit.bnot(4294901759)
    require:NoError(t, err)
    assert:Equal(t, 65536, result)
end
