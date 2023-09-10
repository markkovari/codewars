import { expect, test, it, describe } from "bun:test";
import { toWeirdCase } from './index'

describe('toWeirdCase', function () {
    test('should return the correct value for a single word', function () {
        expect(toWeirdCase('This')).toEqual('ThIs');
        expect(toWeirdCase('is')).toEqual('Is');
    });
    test('should return the correct value for multiple words', function () {
        expect(toWeirdCase('This is a test')).toEqual('ThIs Is A TeSt');
    });
});